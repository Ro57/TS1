package issueancerpc

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/DB"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/issuer"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/justifications"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/lock"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/encoder"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/tokendb"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/utils"
	"github.com/pkt-cash/pktd/lnd/lnwallet"
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
)

type IssunceEvents struct {
	StopSig chan struct{}
}

type Server struct {
	// Nest unimplemented server implementation in order to satisfy server interface
	events          IssunceEvents
	Client          replicator.ReplicatorClient
	cfg             *Config
	chain           lnwallet.BlockChainIO
	db              *tokendb.TokenStrikeDB
	clientSyncChain replicator.Replicator_SyncChainClient
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

func RunServerServing(host string, replicationHost string, events IssunceEvents, db *tokendb.TokenStrikeDB, chain lnwallet.BlockChainIO) {
	client, closeConn, err := connectReplicatorClient(context.TODO(), replicationHost)
	if err != nil {
		panic(err)
	}

	var (
		child = &Server{
			Client: client,
			events: events,
			chain:  chain,
			db:     db,
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

func (s *Server) GetSyncChainClient(ctx context.Context) (replicator.Replicator_SyncChainClient, error) {
	if s.clientSyncChain == nil {
		var err error
		s.clientSyncChain, err = s.Client.SyncChain(ctx)
		if err != nil {
			return nil, err
		}
	}
	return s.clientSyncChain, nil
}

func (s *Server) SyncBlock(ctx context.Context, name string, blocks ...*DB.Block) error {
	request := &replicator.SyncChainRequest{
		Name:   name,
		Blocks: blocks,
	}

	client, err := s.GetSyncChainClient(ctx)
	if err != nil {
		return err
	}

	return client.Send(request)
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
	var err error

	//get url for token
	reqGetUrl := &replicator.GetUrlSequenceRequest{
		TokenName: req.Name,
	}
	respGetUrl, errGetUrl := s.Client.GetUrlSequence(ctx, reqGetUrl)
	if errGetUrl != nil {
		return nil, errGetUrl
	}

	if req.Offer.Urls == nil {
		req.Offer.Urls = make([]string, 0, 1)
	}

	req.Offer.Urls = append(req.Offer.Urls, respGetUrl.Url)

	// issuing token
	errIssueToken := s.issueTokenDB(req)
	if errIssueToken != nil {
		return nil, errIssueToken.Native()
	}

	// generate genesis block
	req.Block, err = s.generateGenesisBlock(req.Name)
	if err != nil {
		return nil, err
	}

	// duplicate in the replicator
	_, err = s.Client.IssueToken(ctx, req)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) LockToken(ctx context.Context, req *issuer.LockTokenRequest) (*issuer.LockTokenResponse, error) {
	// lock token
	errLockToken := s.lockTokenDB(req)
	if errLockToken != nil {
		return nil, errLockToken.Native()
	}

	// generate lock block
	lockBlock, err := s.generateLockBlock(req)
	if err != nil {
		return nil, err
	}

	// send block to the replicator
	err = s.SyncBlock(ctx, req.Token, lockBlock)
	if err != nil {
		return nil, err
	}

	lockID, err := s.generateLockID(lockBlock.Justification.(*DB.Block_Lock))
	if err != nil {
		return nil, err
	}

	resp := &issuer.LockTokenResponse{
		LockId: *lockID,
	}

	return resp, nil
}

func (s *Server) GetTokenList(ctx context.Context, req *replicator.GetTokenListRequest) (*replicator.GetTokenListResponse, error) {
	resultList, err := utils.GetTokenList(s.db)
	if err != nil {
		return nil, err
	}

	return &replicator.GetTokenListResponse{
		Tokens: resultList,
		Total:  int32(len(resultList)),
	}, nil
}

func connectReplicatorClient(ctx context.Context, replicationHost string) (_ replicator.ReplicatorClient, closeConn func() error, _ error) {
	if replicationHost == "" {
		return nil, nil, utils.EmptyAddressErr.Native()
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

func (s *Server) issueTokenDB(req *replicator.IssueTokenRequest) er.R {
	return s.db.Update(func(tx walletdb.ReadWriteTx) er.R {
		rootBucket, err := tx.CreateTopLevelBucket(utils.TokensKey)
		if err != nil {
			return err
		}

		tokenBucket, err := rootBucket.CreateBucket([]byte(req.Name))
		if err != nil {
			return err
		}

		// if information about token did not exist then create
		if tokenBucket.Get(utils.InfoKey) == nil {
			tokenBytes, err := proto.Marshal(req.Offer)
			if err != nil {
				return er.E(err)
			}

			errPut := tokenBucket.Put(utils.InfoKey, tokenBytes)
			if errPut != nil {
				return errPut
			}
		}

		// if token state did not exist then create
		if tokenBucket.Get(utils.StateKey) == nil {
			state := &DB.State{
				Token:  req.Offer,
				Owners: nil,
				Locks:  nil,
			}

			stateBytes, err := proto.Marshal(state)
			if err != nil {
				return er.E(err)
			}

			errPut := tokenBucket.Put(utils.StateKey, stateBytes)
			if errPut != nil {
				return errPut
			}
		}

		return nil
	})
}

func (s *Server) lockTokenDB(req *issuer.LockTokenRequest) er.R {
	return s.db.Update(func(tx walletdb.ReadWriteTx) er.R {
		var block DB.Block
		var oldState DB.State

		rootBucket := tx.ReadWriteBucket(utils.TokensKey)
		if rootBucket == nil {
			return er.Errorf("%s bucket not created", string(utils.TokensKey))
		}

		tokenBucket := rootBucket.NestedReadWriteBucket([]byte(req.Token))

		// if information about token did not exist then create
		if tokenBucket.Get(utils.InfoKey) == nil {
			return er.Errorf("token %s without information", req.Token)
		}

		blockHash := tokenBucket.Get(utils.RootHashKey)
		if blockHash == nil {
			return utils.LastBlockNotFoundErr
		}

		chainBucket := tokenBucket.NestedReadWriteBucket(utils.ChainKey)
		if chainBucket == nil {
			return utils.ChainBucketNotFoundErr
		}

		jsonBlock := chainBucket.Get(blockHash)
		if blockHash == nil {
			return utils.LastBlockNotFoundErr
		}

		err := proto.Unmarshal(jsonBlock, &block)
		if err != nil {
			return er.Errorf("unmarshal block %s \n form json: %v", jsonBlock, err)
		}

		jsonState := tokenBucket.Get(utils.StateKey)
		if jsonState == nil {
			return utils.StateNotFoundErr
		}

		err = proto.Unmarshal(jsonState, &oldState)
		if err != nil {
			return er.Errorf("unmarshal state form json: %v", err)
		}

		state := &DB.State{
			Token:  oldState.Token,
			Owners: oldState.Owners,
			Locks: append(oldState.Locks, &lock.Lock{
				Count:     req.GetCount(),
				Recipient: req.GetRecipient(),
				// TODO: change on real issuer wallet address
				Sender:         "plk1sender",
				HtlcSecretHash: req.GetHtlc(),
				// TODO: get proof count from config
				ProofCount:     1000,
				CreationHeight: block.Height + 1,
			}),
		}

		stateBytes, err := proto.Marshal(state)
		if err != nil {
			return er.Errorf("marshal new state: %v", err)
		}

		errPut := tokenBucket.Put(utils.StateKey, stateBytes)
		if errPut != nil {
			return errPut
		}

		return nil
	})
}

func (s *Server) generateLockID(blockLock *DB.Block_Lock) (*string, error) {
	lock := blockLock.Lock.Lock

	jsonLock, err := proto.Marshal(lock)
	if err != nil {
		return nil, err
	}

	hashLock := encoder.CreateHash(jsonLock)
	lockID := hex.EncodeToString(hashLock)
	return &lockID, nil
}

func (s *Server) generateLockBlock(req *issuer.LockTokenRequest) (*DB.Block, error) {
	var state DB.State
	var block DB.Block

	lockBlock := &DB.Block{
		Justification: &DB.Block_Lock{},
		Creation:      time.Now().Unix(),
	}

	currentBlockHash, currentBlockHeight, err := s.chain.GetBestBlock()
	if err != nil {
		return nil, err.Native()
	}

	lockBlock.PktBlockHash = currentBlockHash.String()
	lockBlock.PktBlockHeight = currentBlockHeight

	err = s.db.Update(func(tx walletdb.ReadWriteTx) er.R {
		rootBucket, err := tx.CreateTopLevelBucket(utils.TokensKey)
		if err != nil {
			return err
		}

		tokenBucket := rootBucket.NestedReadWriteBucket([]byte(req.Token))

		lastHash := tokenBucket.Get(utils.RootHashKey)
		if lastHash == nil {
			return utils.LastBlockNotFoundErr
		}

		chainBucket := tokenBucket.NestedReadWriteBucket(utils.ChainKey)
		if chainBucket == nil {
			return utils.ChainBucketNotFoundErr
		}

		jsonBlock := chainBucket.Get(lastHash)
		if jsonBlock == nil {
			return utils.LastBlockNotFoundErr
		}

		nativeErr := proto.Unmarshal(jsonBlock, &block)
		if nativeErr != nil {
			return er.Errorf("unmarshal block form json: %v", nativeErr)
		}

		jsonState := tokenBucket.Get(utils.StateKey)
		if jsonState == nil {
			return utils.StateNotFoundErr
		}

		nativeErr = proto.Unmarshal(jsonState, &state)
		if nativeErr != nil {
			return er.Errorf("marshal new state: %v", nativeErr)
		}

		lock := state.Locks[len(state.Locks)-1]
		lockBlock.Justification = &DB.Block_Lock{Lock: &justifications.LockToken{Lock: lock}}
		lockBlock.Height = block.Height + 1

		hashState := encoder.CreateHash(jsonState)
		lockBlock.State = hex.EncodeToString(hashState)
		// TODO: Change to signature generation
		lockBlock.Signature = lockBlock.GetState()
		lockBlock.PrevBlock = string(lastHash)

		lockBlockBytes, nativeErr := proto.Marshal(lockBlock)
		if nativeErr != nil {
			return err
		}

		blockSignatureBytes := []byte(lockBlock.GetSignature())
		err = tokenBucket.Put(utils.RootHashKey, blockSignatureBytes)
		if err != nil {
			return err
		}

		return chainBucket.Put(blockSignatureBytes, lockBlockBytes)
	})

	if err != nil {
		return nil, err.Native()
	}

	return lockBlock, nil
}

func (s *Server) generateGenesisBlock(name string) (*DB.Block, error) {
	genesisBlock := &DB.Block{
		PrevBlock:      "",
		Justification:  &DB.Block_Genesis{Genesis: &justifications.Genesis{Token: name}}, //TODO: setup the justification
		Creation:       time.Now().Unix(),
		State:          "",
		PktBlockHash:   "",
		PktBlockHeight: 0,
		Height:         0,
		Signature:      "",
	}

	currentBlockHash, currentBlockHeight, err := s.chain.GetBestBlock()
	if err != nil {
		return nil, err.Native()
	}

	genesisBlock.PktBlockHash = currentBlockHash.String()
	genesisBlock.PktBlockHeight = currentBlockHeight

	err = s.db.Update(func(tx walletdb.ReadWriteTx) er.R {
		rootBucket, err := tx.CreateTopLevelBucket(utils.TokensKey)
		if err != nil {
			return err
		}

		tokenBucket := rootBucket.NestedReadWriteBucket([]byte(name))

		state := tokenBucket.Get(utils.StateKey)
		if state == nil {
			return utils.StateNotFoundErr
		}

		hashState := encoder.CreateHash(state)
		genesisBlock.State = hex.EncodeToString(hashState)

		_, errMarshal := proto.Marshal(genesisBlock)
		if errMarshal != nil {
			return err
		}

		genesisBlock.Signature = genesisBlock.State // TODO: setup the signature

		log.Infof("block root fo issuer %s", genesisBlock.GetSignature())

		blockSignatureBytes := []byte(genesisBlock.GetSignature())
		err = tokenBucket.Put(utils.RootHashKey, blockSignatureBytes)
		if err != nil {
			return err
		}

		genBlockBytes, errMarshal := proto.Marshal(genesisBlock)
		if errMarshal != nil {
			return er.E(errMarshal)
		}

		chainBucket, err := tokenBucket.CreateBucketIfNotExists(utils.ChainKey)
		if err != nil {
			return err
		}

		return chainBucket.Put(blockSignatureBytes, genBlockBytes)
	})
	if err != nil {
		return nil, err.Native()
	}

	return genesisBlock, nil
}
