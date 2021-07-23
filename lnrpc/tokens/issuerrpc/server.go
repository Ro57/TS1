package issueancerpc

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"sync"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/lightninglabs/protobuf-hex-display/json"
	"github.com/pkg/errors"
	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/issuer"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/encoder"
	"github.com/pkt-cash/pktd/lnd/macaroons"
	"github.com/pkt-cash/pktd/pktlog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/macaroon-bakery.v2/bakery"
)

const (
	subServerName = "IssuanceRPC"
)

var (
	tokens sync.Map

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
)

type token struct {
	price       uint64
	validTime   int64
	issuerLogin string
}

// RevokeSig — on-chain addresses for sending coins after token revoke
type RevokeSig struct {
	Token        string
	AddrToAmount map[string]int64
}

type IssunceEvents struct {
	StopSig     chan struct{}
	RevokeEvent chan RevokeSig
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

// Override method of unimplemented server
func (s *Server) SignTokenPurchase(ctx context.Context, req *issuer.SignTokenPurchaseRequest) (*issuer.SignTokenPurchaseResponse, error) {
	bytes, err := json.Marshal(req)
	if err != nil {
		return nil, errors.WithMessage(err, "marshalling request")
	}

	hash := sha256.Sum256(bytes)

	resp := &issuer.SignTokenPurchaseResponse{
		IssuerSignature: fmt.Sprintf("%x", hash),
	}

	return resp, nil
}

func (s *Server) SignTokenSell(ctx context.Context, req *issuer.SignTokenSellRequest) (*issuer.SignTokenSellResponse, error) {
	tokenSell := encoder.TokenSell{
		Token:            req.Offer.Token,
		Price:            req.Offer.Price,
		ID:               req.Offer.IssuerInfo.Id,
		Identity_pubkey:  req.Offer.IssuerInfo.IdentityPubkey,
		Host:             req.Offer.IssuerInfo.Host,
		TokenHolderLogin: req.Offer.TokenHolderLogin,
		TokenBuyerLogin:  req.Offer.TokenBuyerLogin,
		ValidUntilTime:   req.Offer.ValidUntilSeconds,
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

func (s *Server) IssueToken(ctx context.Context, req *issuer.IssueTokenRequest) (*empty.Empty, error) {
	replicatorReq := &replicator.VerifyIssuerRequest{
		Login: req.Offer.TokenHolderLogin,
	}

	_, err := s.Client.VerifyIssuer(ctx, replicatorReq)

	if err != nil {
		return nil, err
	}

	_, ok := tokens.Load(req.Offer.Token)
	if ok {
		return nil, status.Error(codes.InvalidArgument, "token with this name already exists")
	}

	tokens.Store(req.Offer.Token, token{
		issuerLogin: req.Offer.TokenHolderLogin,
		price:       req.Offer.Price,
		validTime:   req.Offer.ValidUntilSeconds,
	})

	return &emptypb.Empty{}, nil
}

func (s *Server) UpdateToken(ctx context.Context, req *issuer.UpdateTokenRequest) (*empty.Empty, error) {
	replicatorReq := &replicator.VerifyIssuerRequest{
		Login: req.Offer.TokenHolderLogin,
	}

	_, err := s.Client.VerifyIssuer(ctx, replicatorReq)
	if err != nil {
		return nil, err
	}

	t, ok := tokens.Load(req.Offer.Token)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "token with this name does not exist")
	}

	newToken := t.(token)

	if req.Offer.Price != 0 {
		newToken.price = req.Offer.Price
	}

	if req.Offer.ValidUntilSeconds != 0 {
		newToken.validTime = req.Offer.ValidUntilSeconds
	}

	tokens.Store(req.Offer.Token, newToken)

	return &emptypb.Empty{}, nil
}

func (s *Server) RevokeToken(ctx context.Context, req *issuer.RevokeTokenRequest) (*empty.Empty, error) {
	replicatorReq := &replicator.VerifyIssuerRequest{
		Login: req.Login,
	}

	_, err := s.Client.VerifyIssuer(ctx, replicatorReq)
	if err != nil {
		return nil, err
	}

	_, ok := tokens.Load(req.TokenName)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "token with this name does not exist")
	}

	amountToAddr := s.payoutCalculate()

	sig := RevokeSig{
		Token:        req.TokenName,
		AddrToAmount: amountToAddr,
	}

	s.events.RevokeEvent <- sig
	tokens.Delete(req.TokenName)

	return &emptypb.Empty{}, nil
}

func (s *Server) GetTokenList(ctx context.Context, req *issuer.GetTokenListRequest) (*issuer.GetTokenListResponse, error) {
	var tokenList []*replicator.TokenOffer
	var err error

	if req.IssuerId != "" {
		tokenList, err = s.issuerTokens(ctx, req.IssuerId)
	} else {
		tokenList, err = s.allTokens()
	}

	if err != nil {
		return nil, err
	}

	return &issuer.GetTokenListResponse{
		Tokens: tokenList,
		Total:  int32(len(tokenList)),
	}, nil

}

func (s *Server) issuerTokens(ctx context.Context, login string) ([]*replicator.TokenOffer, error) {
	resultList := []*replicator.TokenOffer{}

	replicatorReq := &replicator.VerifyIssuerRequest{
		Login: login,
	}

	_, err := s.Client.VerifyIssuer(ctx, replicatorReq)

	if err != nil {
		return nil, err
	}

	tokens.Range(func(key, value interface{}) bool {
		t := value.(token)

		if t.issuerLogin == login {
			resultList = append(resultList, &replicator.TokenOffer{
				Token:             key.(string),
				Price:             t.price,
				ValidUntilSeconds: t.validTime,
				IssuerInfo: &replicator.IssuerInfo{
					Id: login,
				},
			})
		}

		return true
	})

	return resultList, nil
}

func (s *Server) allTokens() ([]*replicator.TokenOffer, error) {
	resultList := []*replicator.TokenOffer{}

	tokens.Range(func(key, value interface{}) bool {
		t := value.(token)

		resultList = append(resultList, &replicator.TokenOffer{
			Token:             key.(string),
			Price:             t.price,
			ValidUntilSeconds: t.validTime,
			IssuerInfo: &replicator.IssuerInfo{
				Id: t.issuerLogin,
			},
		})

		return true
	})

	return resultList, nil
}

// TODO: implement colculation of real amount of pkt and send on real addresses
func (s *Server) payoutCalculate() map[string]int64 {
	return map[string]int64{
		"alice": 1000,
		"bob":   2000,
	}
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
