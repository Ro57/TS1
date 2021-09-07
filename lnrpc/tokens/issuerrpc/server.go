package issueancerpc

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/lightninglabs/protobuf-hex-display/json"
	"github.com/pkg/errors"
	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/issuer"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/encoder"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/tokendb"
	"github.com/pkt-cash/pktd/lnd/macaroons"
	"github.com/pkt-cash/pktd/pktlog/log"
	"github.com/pkt-cash/pktd/pktwallet/walletdb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/macaroon-bakery.v2/bakery"
)

const (
	subServerName = "IssuanceRPC"
)

var (

	// macaroonOps are the set of capabilities that our minted macaroon (if
	// it doesn't already exist) will have.
	macaroonOps = []bakery.Op{
		{
			Entity: "issuance",
			Action: "read",
		},
	}

	// macPermissions maps RPC calls to the permissions they require.
	macPermissions = map[string][]bakery.Op{
		"/issuancerpc.Issuance/SignTokenSell": {{
			Entity: "issuance",
			Action: "read",
		}},
	}

	// DefaultIssuanceMacFilename is the default name of the issuance macaroon
	// that we expect to find via a file handle within the main
	// configuration file in this package.
	DefaultIssuanceMacFilename = "issuance.macaroon"

	infoKey     = []byte("info")
	chainKey    = []byte("chain")
	tokensKey   = []byte("tokens")
	rootHashKey = []byte("rootHash")
)

type IssunceEvents struct {
	StopSig chan struct{}
}

type Server struct {
	// Nest unimplemented server implementation in order to satisfy server interface
	events IssunceEvents
	Client replicator.ReplicatorClient
	cfg    *Config
}

// New returns a new instance of the issueancerpc Issuer sub-server. We also return
// the set of permissions for the macaroons that we may create within this
// method. If the macaroons we need aren't found in the filepath, then we'll
// create them on start up. If we're unable to locate, or create the macaroons
// we need, then we'll return with an error.
func New(cfg *Config) (*Server, lnrpc.MacaroonPerms, er.R) {
	// If the path of the issuance macaroon wasn't generated, then we'll
	// assume that it's found at the default network directory.
	if cfg.IssuanceMacPath == "" {
		cfg.IssuanceMacPath = filepath.Join(
			cfg.NetworkDir, DefaultIssuanceMacFilename,
		)
	}

	// Now that we know the full path of the signer macaroon, we can check
	// to see if we need to create it or not. If stateless_init is set
	// then we don't write the macaroons.
	macFilePath := cfg.IssuanceMacPath
	if cfg.MacService != nil && !cfg.MacService.StatelessInit &&
		!lnrpc.FileExists(macFilePath) {

		log.Infof("Making macaroons for Issuance RPC Server at: %v",
			macFilePath)

		// At this point, we know that the signer macaroon doesn't yet,
		// exist, so we need to create it with the help of the main
		// macaroon service.
		IssuanceMac, err := cfg.MacService.NewMacaroon(
			context.Background(), macaroons.DefaultRootKeyID,
			macaroonOps...,
		)
		if err != nil {
			return nil, nil, err
		}
		signerMacBytes, errr := IssuanceMac.M().MarshalBinary()
		if errr != nil {
			return nil, nil, er.E(errr)
		}
		errr = ioutil.WriteFile(macFilePath, signerMacBytes, 0644)
		if errr != nil {
			_ = os.Remove(macFilePath)
			return nil, nil, er.E(errr)
		}
	}

	issuanceServer := &Server{
		cfg: cfg,
	}

	return issuanceServer, macPermissions, nil
}

func RunServerServing(host string, replicationHost string, events IssunceEvents) {
	client, closeConn, err := connectReplicatorClient(context.TODO(), replicationHost)
	if err != nil {
		panic(err)
	}

	var (
		child = &Server{
			Client: client,
			events: events,
		}
		root = grpc.NewServer()
	)
	issuer.RegisterIssuerServiceServer(root, child)

	listener, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}

	go func() {
		err := root.Serve(listener)
		if err != nil {
			panic(err)
		}
		closeConn()
	}()

	go func() {
		<-events.StopSig
		root.Stop()
		closeConn()
	}()
}

// Start launches any helper goroutines required for the rpcServer to function.
//
// NOTE: This is part of the lnrpc.SubServer interface.
func (s *Server) Start() er.R {
	return nil
}

// Stop signals any active goroutines for a graceful closure.
//
// NOTE: This is part of the lnrpc.SubServer interface.
func (s *Server) Stop() er.R {
	return nil
}

// Name returns a unique string representation of the sub-server. This can be
// used to identify the sub-server and also de-duplicate them.
//
// NOTE: This is part of the lnrpc.SubServer interface.
func (s *Server) Name() string {
	return subServerName
}

// RegisterWithRootServer will be called by the root gRPC server to direct a
// sub RPC server to register itself with the main gRPC root server. Until this
// is called, each sub-server won't be able to have
// requests routed towards it.
//
// NOTE: This is part of the lnrpc.SubServer interface.
func (s *Server) RegisterWithRootServer(grpcServer *grpc.Server) er.R {
	// We make sure that we register it with the main gRPC server to ensure
	// all our methods are routed properly.
	issuer.RegisterIssuerServiceServer(grpcServer, s)

	log.Debugf("Issuance RPC server successfully register with root gRPC " +
		"server")

	return nil
}

// RegisterWithRestServer will be called by the root REST mux to direct a sub
// RPC server to register itself with the main REST mux server. Until this is
// called, each sub-server won't be able to have requests routed towards it.
//
// NOTE: This is part of the lnrpc.SubServer interface.
func (s *Server) RegisterWithRestServer(ctx context.Context,
	mux *runtime.ServeMux, dest string, opts []grpc.DialOption) er.R {
	// TODO: Clarify whether it is necessary REST API, and if it necessary
	// describe rest notation in yaml file and generate .gw file from proto
	// notation. Implementation of RegisterWithRestServer can be found in
	// other services, such as the signature service.
	return nil
}

func (s *Server) SignTokenSell(ctx context.Context, req *issuer.SignTokenSellRequest) (*issuer.SignTokenSellResponse, error) {
	tokenSell := encoder.TokenSell{
		Token:          req.Offer.Token,
		Price:          req.Offer.Price,
		ValidUntilTime: req.Offer.ValidUntilSeconds,
		Count:          req.Offer.Count,
	}

	bytes, err := json.Marshal(tokenSell)
	if err != nil {
		return nil, errors.WithMessage(err, "marshalling request")
	}

	hash := encoder.CreateHash(bytes)

	resp := &issuer.SignTokenSellResponse{
		IssuerSignature: fmt.Sprintf("%x", hash),
	}

	return resp, nil
}

func (s *Server) IssueToken(ctx context.Context, req *replicator.IssueTokenRequest) (*empty.Empty, error) {
	err := issueTokenDB(req).Native()
	if err != nil {
		return nil, err
	}

	_, err = s.Client.IssueToken(ctx, req)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) GetTokenList(ctx context.Context, req *replicator.GetTokenListRequest) (*replicator.GetTokenListResponse, error) {
	//TODO: not call client, should call DB methods
	return s.Client.GetTokenList(ctx, req)
}

func connectReplicatorClient(ctx context.Context, replicationHost string) (_ replicator.ReplicatorClient, closeConn func() error, _ error) {
	if replicationHost == "" {
		return nil, nil, errors.New("empty address")
	}

	// TODO: research connection option to be secure for protected methods
	// 	? Use "r.restDialOpts"
	conn, err := grpc.DialContext(
		ctx,
		replicationHost,
		grpc.WithInsecure(),
	)

	if err != nil {
		return nil, nil, errors.WithMessage(err, "dialing")
	}
	return replicator.NewReplicatorClient(conn), conn.Close, nil
}

func issueTokenDB(req *replicator.IssueTokenRequest) er.R {
	return tokendb.Update(func(tx walletdb.ReadWriteTx) er.R {
		rootBucket, err := tx.CreateTopLevelBucket(tokensKey)
		if err != nil {
			return err
		}

		tokenBucket, err := rootBucket.CreateBucket([]byte(req.Name))
		if err != nil {
			return err
		}

		// if information about token did not exist then create
		if tokenBucket.Get(infoKey) == nil {
			tokenBytes, err := json.Marshal(req.Offer)
			if err != nil {
				return er.E(err)
			}

			errPut := tokenBucket.Put(infoKey, tokenBytes)
			if errPut != nil {
				return errPut
			}
		}

		return nil
	})
}
