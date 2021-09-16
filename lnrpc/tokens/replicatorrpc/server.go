package replicatorrpc

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/DB"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/jwtstore"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/tokendb"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/utils"
	"github.com/pkt-cash/pktd/lnd/lnwallet"
	"github.com/pkt-cash/pktd/lnd/macaroons"
	"github.com/pkt-cash/pktd/pktlog/log"
	"github.com/pkt-cash/pktd/pktwallet/walletdb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/macaroon-bakery.v2/bakery"
)

const (
	subServerName = "ReplicatorRPC"
)

// token holders with login password
var (
	jwtStore *jwtstore.Store
	tokens   sync.Map

	signingKey = []byte("SUPER_SECRET")

	// macaroonOps are the set of capabilities that our minted macaroon (if
	// it doesn't already exist) will have.
	macaroonOps = []bakery.Op{
		{
			Entity: "replicator",
			Action: "read",
		},
	}

	// macPermissions maps RPC calls to the permissions they require.
	macPermissions = map[string][]bakery.Op{
		"/replicatorrpc.Replicator/GetTokenBalances": {{
			Entity: "replicator",
			Action: "read",
		}},
	}

	// DefaultReplicatorMacFilename is the default name of the replicator macaroon
	// that we expect to find via a file handle within the main
	// configuration file in this package.
	DefaultReplicatorMacFilename = "replicator.macaroon"
)

type userInfo struct {
	password string
	// TODO: implement role system
	// ? method for add new role to user
	roles    map[string]struct{}
	balances TokenHolderBalancesStore
}

type OpenChannel struct {
	Address lnrpc.LightningAddress
	Amount  int64
}

type ReplicatorEvents struct {
	StopSig          chan struct{}
	OpenChannelEvent chan OpenChannel
	RevokeEvent      chan RevokeSig
	OpenChannelError chan error
}

// RevokeSig — on-chain addresses for sending coins after token revoke
type RevokeSig struct {
	Token        string
	AddrToAmount map[string]int64
}

type Server struct {
	// Nest unimplemented server implementation in order to satisfy server interface
	replicator.UnimplementedReplicatorServer

	users sync.Map

	events ReplicatorEvents
	cfg    *Config
	chain  lnwallet.BlockChainIO
	db     *tokendb.TokenStrikeDB
}

type loginCliams struct {
	login string
	jwt.StandardClaims
}

type token struct {
	price     uint64
	validTime int64
}

var _ replicator.ReplicatorServer = (*Server)(nil)

// New returns a new instance of the replicatorrpc Repicator sub-server. We also return
// the set of permissions for the macaroons that we may create within this
// method. If the macaroons we need aren't found in the filepath, then we'll
// create them on start up. If we're unable to locate, or create the macaroons
// we need, then we'll return with an error.
func New(cfg *Config) (*Server, lnrpc.MacaroonPerms, er.R) {
	// If the path of the replicator macaroon wasn't generated, then we'll
	// assume that it's found at the default network directory.
	if cfg.ReplicatorMacPath == "" {
		cfg.ReplicatorMacPath = filepath.Join(
			cfg.NetworkDir, DefaultReplicatorMacFilename,
		)
	}

	// Now that we know the full path of the replicator macaroon, we can check
	// to see if we need to create it or not. If stateless_init is set
	// then we don't write the macaroons.
	macFilePath := cfg.ReplicatorMacPath
	if cfg.MacService != nil && !cfg.MacService.StatelessInit &&
		!lnrpc.FileExists(macFilePath) {

		log.Infof("Making macaroons for replicator RPC Server at: %v",
			macFilePath)

		// At this point, we know that the replicator macaroon doesn't yet,
		// exist, so we need to create it with the help of the main
		// macaroon service.
		replicatorMac, err := cfg.MacService.NewMacaroon(
			context.Background(), macaroons.DefaultRootKeyID,
			macaroonOps...,
		)

		if err != nil {
			return nil, nil, err
		}
		replicatorMacBytes, errr := replicatorMac.M().MarshalBinary()
		if errr != nil {
			return nil, nil, er.E(errr)
		}
		errr = ioutil.WriteFile(macFilePath, replicatorMacBytes, 0644)
		if errr != nil {
			_ = os.Remove(macFilePath)
			return nil, nil, er.E(errr)
		}
	}

	replicatorServer := &Server{
		cfg: cfg,
	}

	return replicatorServer, macPermissions, nil

}
func RunServerServing(host string, events ReplicatorEvents, db *tokendb.TokenStrikeDB, chain lnwallet.BlockChainIO) {

	var (
		child = &Server{
			events: events,
			chain:  chain,
			db:     db,
		}
		root = grpc.NewServer()
	)
	replicator.RegisterReplicatorServer(root, child)

	listener, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}

	go func() {
		err := root.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		<-events.StopSig
		root.Stop()
	}()

	jwtStore = jwtstore.New([]jwtstore.JWT{})
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
	replicator.RegisterReplicatorServer(grpcServer, s)

	log.Debugf("Replicator RPC server successfully register with root gRPC " +
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
func (s *Server) GetTokenOffers(ctx context.Context, req *replicator.GetTokenOffersRequest) (*replicator.GetTokenOffersResponse, error) {
	const (
		eachIssuerTokensNum = 3
		offersNum           = 1000 * eachIssuerTokensNum
	)

	offers := make([]*replicator.TokenOffer, 0, offersNum)

	// Fill mocked offers such, that each issuer has several tokens present
	for i := offersNum / eachIssuerTokensNum; i > 0; i-- {
		offer := &replicator.TokenOffer{
			ValidUntilSeconds: time.Now().Unix() + int64(i)*1000,
			IssuerInfo: &replicator.IssuerInfo{
				Id:             fmt.Sprintf("issuer_%d", i),
				IdentityPubkey: "issuer_node_pub_key",
				Host:           "issuer_ip",
			},
			Token: fmt.Sprintf("token_%d", i),
			Price: uint64(1 + i*2),
		}
		offers = append(offers, offer)

		offer = &replicator.TokenOffer{
			ValidUntilSeconds: time.Now().Unix() + int64(i)*1000,
			IssuerInfo: &replicator.IssuerInfo{
				Id:             fmt.Sprintf("issuer_%d", i),
				IdentityPubkey: "issuer_node_pub_key",
				Host:           "issuer_ip",
			},
			Token: fmt.Sprintf("token_%d", i+1),
			Price: uint64(1 + i*4),
		}
		offers = append(offers, offer)

		offer = &replicator.TokenOffer{
			ValidUntilSeconds: time.Now().Unix() + int64(i)*1000,
			IssuerInfo: &replicator.IssuerInfo{
				Id:             fmt.Sprintf("issuer_%d", i),
				IdentityPubkey: "issuer_node_pub_key",
				Host:           "issuer_ip",
			},
			Token: fmt.Sprintf("token_%d", i+2),
			Price: uint64(1 + i*8),
		}
		offers = append(offers, offer)
	}

	resp := &replicator.GetTokenOffersResponse{
		Offers: offers,
		Total:  offersNum,
	}

	// Apply filter by issuer id
	if req.IssuerId != "" {
		issuerOffers := make([]*replicator.TokenOffer, 0, eachIssuerTokensNum)

		for _, offer := range resp.Offers {
			if len(issuerOffers) == eachIssuerTokensNum {
				break
			}

			if offer.IssuerInfo.Id == req.IssuerId {
				issuerOffers = append(issuerOffers, offer)
			}
		}

		resp.Offers = issuerOffers
		resp.Total = uint64(len(resp.Offers))
	}

	// Apply pagination
	if req.Params.Offset > 0 {
		if int(req.Params.Offset) <= len(resp.Offers)-1 {
			resp.Offers = resp.Offers[req.Params.Offset:]
		} else {
			resp.Offers = nil
		}
	}
	if req.Params.Limit > 0 {
		if int(req.Params.Limit) <= len(resp.Offers)-1 {
			resp.Offers = resp.Offers[:req.Params.Limit]
		}
	}

	return resp, nil
}

// Override method of unimplemented server
func (s *Server) GetTokenBalances(ctx context.Context, req *replicator.GetTokenBalancesRequest) (*replicator.GetTokenBalancesResponse, error) {

	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "metadata not provided")
	}

	tokenSet, ok := meta["jwt"]
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "jwt token is not contained in context")
	}

	// Get first token. By default MD contain slice of strings
	// But we need only one jwt
	tokenHash := tokenSet[0]

	innerJWT, err := jwtStore.GetByToken(tokenHash)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("jwt by token not found: %v", err))
	}

	if innerJWT.ExpireDate.Before(time.Now()) {
		return nil, status.Error(codes.ResourceExhausted, "session expired")
	}

	v, ok := s.users.Load(innerJWT.HolderLogin)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("user with login %v does not exist ", innerJWT.HolderLogin))
	}

	info := v.(userInfo)

	resp := &replicator.GetTokenBalancesResponse{
		Balances: info.balances.store,
		Total:    uint64(len(info.balances.store)),
	}

	return resp, nil
}

func (s *Server) VerifyIssuer(ctx context.Context, req *replicator.VerifyIssuerRequest) (*empty.Empty, error) {
	if req.Login == "" {
		return nil, status.Error(codes.InvalidArgument, "user login not presented")
	}

	user, ok := s.users.Load(req.Login)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "user with this login does not exist")
	}

	ui := user.(userInfo)

	_, ok = ui.roles["issuer"]
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "this user was not permitted to create new tokens")
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) IssueToken(ctx context.Context, req *replicator.IssueTokenRequest) (*empty.Empty, error) {
	err := s.db.Update(func(tx walletdb.ReadWriteTx) er.R {
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
			state := DB.State{
				Token:  req.Offer,
				Owners: nil,
				Locks:  nil,
			}

			stateBytes, err := proto.Marshal(&state)
			if err != nil {
				return er.E(err)
			}

			errPut := tokenBucket.Put(utils.StateKey, stateBytes)
			if errPut != nil {
				return errPut
			}
		}

		return tokenBucket.Put(utils.RootHashKey, []byte(""))
	})
	if err != nil {
		return nil, err.Native()
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) GetTokenList(ctx context.Context, req *replicator.GetTokenListRequest) (*replicator.GetTokenListResponse, error) {
	tokenList, err := s.allTokensFromIssuer()
	if err != nil {
		return nil, err
	}

	return &replicator.GetTokenListResponse{
		Tokens: tokenList,
		Total:  int32(len(tokenList)),
	}, nil

}

func (s *Server) SaveBlock(ctx context.Context, req *replicator.SaveBlockRequest) (*empty.Empty, error) {
	err := s.db.Update(func(tx walletdb.ReadWriteTx) er.R {
		rootBucket, err := tx.CreateTopLevelBucket(utils.TokensKey)
		if err != nil {
			return err
		}

		tokenBucket := rootBucket.NestedReadWriteBucket([]byte(req.Name))

		if string(tokenBucket.Get(utils.RootHashKey)) != req.Block.PrevBlock {
			return er.Errorf("invalid hash of the previous block want %s but get %s", tokenBucket.Get(utils.RootHashKey), req.Block.PrevBlock)
		}

		log.Infof("block root fo replication %s", req.Block.GetSignature())

		blockSignatureBytes := []byte(req.Block.GetSignature())

		err = tokenBucket.Put(utils.RootHashKey, blockSignatureBytes)
		if err != nil {
			return err
		}

		blockBytes, errMarshal := proto.Marshal(req.Block)
		if errMarshal != nil {
			return er.E(errMarshal)
		}

		chainBucket, err := tokenBucket.CreateBucketIfNotExists(utils.ChainKey)
		if err != nil {
			return err
		}
		return chainBucket.Put(blockSignatureBytes, blockBytes)
	})
	if err != nil {
		return nil, err.Native()
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) GetToken(ctx context.Context, req *replicator.GetTokenRequest) (*replicator.GetTokenResponse, error) {
	response := &replicator.GetTokenResponse{
		Token: &replicator.Token{
			Name:  req.TokenId,
			Token: &DB.Token{},
			Root:  "",
		},
	}

	err := s.db.View(func(tx walletdb.ReadTx) er.R {
		rootBucket := tx.ReadBucket(utils.TokensKey)
		if rootBucket == nil {
			return er.New("tokens DB not created")
		}

		tokenBucket := rootBucket.NestedReadBucket([]byte(req.TokenId))
		if tokenBucket == nil {
			return er.Errorf("token by name %v not found", req.TokenId)
		}
		infoBytes := tokenBucket.Get(utils.InfoKey)
		if infoBytes == nil {
			return er.New("info does not exist")
		}

		err := proto.Unmarshal(infoBytes, response.Token.Token)
		if err != nil {
			return er.E(err)
		}

		response.Token.Root = string(tokenBucket.Get(utils.RootHashKey))
		return nil
	})
	if err != nil {
		return nil, err.Native()
	}

	return response, nil
}

func (s *Server) GetHeaders(ctx context.Context, req *replicator.GetHeadersRequest) (*replicator.GetHeadersResponse, error) {
	response := &replicator.GetHeadersResponse{
		Token:  &DB.Token{},
		Blocks: []*replicator.MerkleBlock{},
	}
	err := s.db.View(func(tx walletdb.ReadTx) er.R {
		tokensBucket := tx.ReadBucket(utils.TokensKey)
		tokenBucket := tokensBucket.NestedReadBucket([]byte(req.TokenId))
		if tokenBucket == nil {
			return er.New("token does not exist")
		}

		infoBytes := tokenBucket.Get(utils.InfoKey)
		if infoBytes == nil {
			return er.New("info does not exist")
		}

		err := proto.Unmarshal(infoBytes, response.Token)
		if err != nil {
			return er.E(err)
		}

		rootHash := tokenBucket.Get(utils.RootHashKey)
		if rootHash == nil {
			return er.New("root hash does not exist exist")
		}

		if string(rootHash) != req.Hash {
			response.Blocks, err = s.getMerkleRoot(tokenBucket, rootHash, req.Hash)
			if err != nil {
				return er.E(err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err.Native()
	}

	return response, nil
}

// TODO: Rework this method. Need geting all issuers and their tokens with wallet addresses
func (s *Server) allTokensFromIssuer() ([]*replicator.Token, error) {
	resultList, err := utils.GetTokenList(s.db)
	if err != nil {
		return nil, err
	}

	return resultList, nil
}

// TODO: implement colculation of real amount of pkt and send on real addresses
func (s *Server) payoutCalculate() map[string]int64 {
	return map[string]int64{
		"alice": 1000,
		"bob":   2000,
	}
}

func (s *Server) getMerkleRoot(bucket walletdb.ReadBucket, root []byte, hash string) ([]*replicator.MerkleBlock, error) {
	var (
		currentHash []byte = root
		response    []*replicator.MerkleBlock
		chainBucket = bucket.NestedReadBucket(utils.ChainKey)
	)

	for {
		blockBytes := chainBucket.Get(currentHash)
		if blockBytes == nil {
			return nil, errors.New("block does not exist exist")
		}

		var block DB.Block
		err := proto.Unmarshal(blockBytes, &block)
		if err != nil {
			return nil, err
		}

		if string(currentHash) == hash {
			break
		}

		merkleBlock := replicator.MerkleBlock{
			Hash:     string(currentHash),
			PrevHash: block.PrevBlock,
		}

		response = append(response, &merkleBlock)
		currentHash = []byte(merkleBlock.PrevHash)
	}

	return response, nil
}

type TokenHoldersStoreAPI interface {
	Insert(TokenHolder) error
	Has(TokenHolder) bool
}
type TokenHoldersStore struct {
	holders map[TokenHolderLogin]TokenHolder
	mu      sync.RWMutex
}

var _ TokenHoldersStoreAPI = (*TokenHoldersStore)(nil)

func NewTokenHoldersStore() *TokenHoldersStore {
	return &TokenHoldersStore{
		holders: make(map[TokenHolderLogin]TokenHolder),
	}
}

func (s *TokenHoldersStore) Insert(holder TokenHolder) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.holders[holder.Login]
	if ok {
		return errors.Errorf("login duplication: %q", holder.Login)
	}

	s.holders[holder.Login] = holder

	return nil
}

func (s *TokenHoldersStore) Has(holder TokenHolder) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := s.holders[holder.Login]
	return ok
}

type TokenHolderLogin string
type TokenName = string

type TokenHolder struct {
	Login    TokenHolderLogin
	Password string
}
