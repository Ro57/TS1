package replicatorrpc

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/encoder"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/jwtstore"
	"github.com/pkt-cash/pktd/lnd/macaroons"
	"github.com/pkt-cash/pktd/pktlog/log"
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
}

type loginCliams struct {
	login string
	jwt.StandardClaims
}

type token struct {
	price     uint64
	validTime int64
}

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
func RunServerServing(host string, events ReplicatorEvents) {
	var (
		child = &Server{
			events: events,
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

// Override method of unimplemented server
func (s *Server) VerifyTokenPurchase(ctx context.Context, req *replicator.VerifyTokenPurchaseRequest) (*empty.Empty, error) {
	// NOTE: is expected to be empty
	if req.Purchase.InitialTxHash != "" {
		return nil, status.Error(codes.InvalidArgument, "initial tx hash is provided")
	}

	if req.Purchase.IssuerSignature == "" {
		return nil, status.Error(codes.InvalidArgument, "issuer signature not provided")
	}

	if req.Purchase.Offer == nil {
		return nil, status.Error(codes.InvalidArgument, "offer's not provided")
	}
	if req.Purchase.Offer.Token == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's token name not provided")
	}
	if req.Purchase.Offer.Price == 0 {
		return nil, status.Error(codes.InvalidArgument, "offer's token price not provided")
	}
	if req.Purchase.Offer.TokenHolderLogin == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's token holder login not provided")
	}
	if req.Purchase.Offer.ValidUntilSeconds == 0 {
		return nil, status.Error(codes.InvalidArgument, "offer's validity until seconds not provided")
	}

	if req.Purchase.Offer.IssuerInfo == nil {
		return nil, status.Error(codes.InvalidArgument, "offer's issuer info not provided")
	}
	if req.Purchase.Offer.IssuerInfo.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's issuer id not provided")
	}
	if req.Purchase.Offer.IssuerInfo.IdentityPubkey == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's issuer identity pubkey not provided")
	}
	if req.Purchase.Offer.IssuerInfo.Host == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's issuer host not provided")
	}

	return &empty.Empty{}, nil
}

func (s *Server) VerifyTokenSell(ctx context.Context, req *replicator.VerifyTokenSellRequest) (*empty.Empty, error) {
	// NOTE: is expected to be empty
	if req.Sell.InitialTxHash != "" {
		return nil, status.Error(codes.InvalidArgument, "initial tx hash is provided")
	}

	if req.Sell.IssuerSignature == "" {
		return nil, status.Error(codes.InvalidArgument, "issuer signature not provided")
	}

	if req.Sell.Offer == nil {
		return nil, status.Error(codes.InvalidArgument, "offer's not provided")
	}
	if req.Sell.Offer.Token == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's token name not provided")
	}
	if req.Sell.Offer.Price == 0 {
		return nil, status.Error(codes.InvalidArgument, "offer's token price not provided")
	}
	if req.Sell.Offer.TokenHolderLogin == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's token holder login not provided")
	}
	if req.Sell.Offer.TokenBuyerLogin == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's token buyer login not provided")
	}
	if req.Sell.Offer.ValidUntilSeconds == 0 {
		return nil, status.Error(codes.InvalidArgument, "offer's validity until seconds not provided")
	}

	if req.Sell.Offer.IssuerInfo == nil {
		return nil, status.Error(codes.InvalidArgument, "offer's issuer info not provided")
	}
	if req.Sell.Offer.IssuerInfo.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's issuer id not provided")
	}
	if req.Sell.Offer.IssuerInfo.IdentityPubkey == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's issuer identity pubkey not provided")
	}
	if req.Sell.Offer.IssuerInfo.Host == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's issuer host not provided")
	}

	tokenSell := encoder.TokenSell{
		Token:            req.Sell.Offer.Token,
		Price:            req.Sell.Offer.Price,
		ID:               req.Sell.Offer.IssuerInfo.Id,
		Identity_pubkey:  req.Sell.Offer.IssuerInfo.IdentityPubkey,
		Host:             req.Sell.Offer.IssuerInfo.Host,
		TokenHolderLogin: req.Sell.Offer.TokenHolderLogin,
		TokenBuyerLogin:  req.Sell.Offer.TokenBuyerLogin,
		ValidUntilTime:   req.Sell.Offer.ValidUntilSeconds,
	}

	bytes, err := json.Marshal(tokenSell)
	if err != nil {
		return nil, errors.WithMessage(err, "marshalling request")
	}

	hash := encoder.CreateHash(bytes)

	if fmt.Sprintf("%x", hash) != req.Sell.IssuerSignature {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprint("Expect hash ", fmt.Sprintf("%x", hash), " but got ", req.Sell.IssuerSignature))
	}

	return &empty.Empty{}, nil
}

func (s *Server) RegisterTokenPurchase(ctx context.Context, req *replicator.RegisterTokenPurchaseRequest) (*empty.Empty, error) {
	login := req.Purchase.Offer.TokenHolderLogin

	if req.Purchase.InitialTxHash == "" {
		return nil, status.Error(codes.InvalidArgument, "initial tx hash not provided")
	}
	if req.Purchase.IssuerSignature == "" {
		return nil, status.Error(codes.InvalidArgument, "issuer signature not provided")
	}

	if req.Purchase.Offer == nil {
		return nil, status.Error(codes.InvalidArgument, "offer's not provided")
	}
	if req.Purchase.Offer.Token == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's token name not provided")
	}
	if req.Purchase.Offer.Price == 0 {
		return nil, status.Error(codes.InvalidArgument, "offer's token price not provided")
	}
	if req.Purchase.Offer.TokenHolderLogin == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's token holder login not provided")
	}
	if req.Purchase.Offer.ValidUntilSeconds == 0 {
		return nil, status.Error(codes.InvalidArgument, "offer's validity until seconds not provided")
	}

	if req.Purchase.Offer.IssuerInfo == nil {
		return nil, status.Error(codes.InvalidArgument, "offer's issuer info not provided")
	}
	if req.Purchase.Offer.IssuerInfo.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's issuer id not provided")
	}
	if req.Purchase.Offer.IssuerInfo.IdentityPubkey == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's issuer identity pubkey not provided")
	}
	if req.Purchase.Offer.IssuerInfo.Host == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's issuer host not provided")
	}

	issuedToken, ok := tokens.Load(req.Purchase.Offer.Token)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("token with name %v was not issued", req.Purchase.Offer.Token))
	}

	issuedTokenInfo := issuedToken.(token)

	PKTPrice := req.Purchase.Offer.Price / issuedTokenInfo.price

	buyer, ok := s.users.Load(login)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("user with login %v does not exist ", login))
	}

	s.events.OpenChannelEvent <- OpenChannel{
		lnrpc.LightningAddress{
			Pubkey: req.Purchase.Offer.IssuerInfo.IdentityPubkey,
			Host:   req.Purchase.Offer.IssuerInfo.Host,
		},

		int64(req.Purchase.Offer.GetPrice()),
	}

	buyerInfo := buyer.(userInfo)

	buyerInfo.balances.AppendOrUpdate([]*replicator.TokenBalance{{
		Token:     req.Purchase.Offer.Token,
		Available: PKTPrice,
		Frozen:    0,
	}})

	return &empty.Empty{}, nil
}

func (s *Server) RegisterTokenSell(ctx context.Context, req *replicator.RegisterTokenSellRequest) (*empty.Empty, error) {
	if req.Sell.InitialTxHash == "" {
		return nil, status.Error(codes.InvalidArgument, "initial tx hash not provided")
	}
	if req.Sell.IssuerSignature == "" {
		return nil, status.Error(codes.InvalidArgument, "issuer signature not provided")
	}

	if req.Sell.Offer == nil {
		return nil, status.Error(codes.InvalidArgument, "offer's not provided")
	}
	if req.Sell.Offer.Token == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's token name not provided")
	}
	if req.Sell.Offer.Price == 0 {
		return nil, status.Error(codes.InvalidArgument, "offer's token price not provided")
	}
	if req.Sell.Offer.TokenHolderLogin == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's token holder login not provided")
	}
	if req.Sell.Offer.TokenBuyerLogin == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's token buyer login not provided")
	}
	if req.Sell.Offer.ValidUntilSeconds == 0 {
		return nil, status.Error(codes.InvalidArgument, "offer's validity until seconds not provided")
	}

	if req.Sell.Offer.IssuerInfo == nil {
		return nil, status.Error(codes.InvalidArgument, "offer's issuer info not provided")
	}
	if req.Sell.Offer.IssuerInfo.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's issuer id not provided")
	}
	if req.Sell.Offer.IssuerInfo.IdentityPubkey == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's issuer identity pubkey not provided")
	}
	if req.Sell.Offer.IssuerInfo.Host == "" {
		return nil, status.Error(codes.InvalidArgument, "offer's issuer host not provided")
	}
	if req.Sell.Offer.TokenHolderLogin == req.Sell.Offer.TokenBuyerLogin {
		return nil, status.Error(codes.InvalidArgument, "buyer and holder don't have to be the same")

	}

	issuedToken, ok := tokens.Load(req.Sell.Offer.Token)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("token with name %v was not issued", req.Sell.Offer.Token))
	}

	holder, ok := s.users.Load(req.Sell.Offer.TokenHolderLogin)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("user with login %v does not exist ", req.Sell.Offer.TokenHolderLogin))
	}

	buyer, ok := s.users.Load(req.Sell.Offer.TokenBuyerLogin)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("user with login %v does not exist ", req.Sell.Offer.TokenBuyerLogin))
	}

	s.events.OpenChannelEvent <- OpenChannel{
		lnrpc.LightningAddress{
			Pubkey: req.Sell.Offer.IssuerInfo.IdentityPubkey,
			Host:   req.Sell.Offer.IssuerInfo.Host,
		},

		int64(req.Sell.Offer.GetPrice()),
	}

	holderInfo := holder.(userInfo)
	buyerInfo := buyer.(userInfo)
	issuedTokenInfo := issuedToken.(token)

	PKTPrice := req.Sell.Offer.Price / issuedTokenInfo.price

	holderBalance, err := holderInfo.balances.Get(req.Sell.Offer.Token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// TODO: Calculate token count. Now we get price of token as 1 PKT
	if holderBalance.Available <= PKTPrice {
		return nil, status.Error(codes.InvalidArgument, "Not enough available funds on the balance sheet")
	}

	// TODO: Calculate token count. Now we get price of token as 1 PKT
	holderInfo.balances.AppendOrUpdate([]*replicator.TokenBalance{{
		Token:     req.Sell.Offer.Token,
		Available: -PKTPrice,
		Frozen:    PKTPrice,
	}})

	// TODO: Check if token contain in store update balance
	buyerInfo.balances.AppendOrUpdate([]*replicator.TokenBalance{{
		Token:     req.Sell.Offer.Token,
		Available: PKTPrice,
		Frozen:    0,
	}})

	s.users.Store(req.Sell.Offer.TokenHolderLogin, holderInfo)
	s.users.Store(req.Sell.Offer.TokenBuyerLogin, buyerInfo)

	// TODO: Replace error hanlder before store a new balance
	err = <-s.events.OpenChannelError

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "registrete sell own token: %v", err)
	}

	return &empty.Empty{}, nil
}

func (s *Server) RegisterTokenHolder(ctx context.Context, req *replicator.RegisterRequest) (*empty.Empty, error) {
	_, ok := s.users.Load(req.Login)
	if ok {
		return nil, status.Error(codes.InvalidArgument, "user with this login already exists")
	}

	roles := make(map[string]struct{})
	roles["holder"] = struct{}{}

	s.users.Store(req.Login, userInfo{
		password: req.Password,
		roles:    roles,
	})

	return &empty.Empty{}, nil
}

func (s *Server) RegisterTokenIssuer(ctx context.Context, req *replicator.RegisterRequest) (*empty.Empty, error) {
	_, ok := s.users.Load(req.Login)
	if ok {
		return nil, status.Error(codes.InvalidArgument, "user with this login already exists")
	}
	roles := make(map[string]struct{})
	roles["issuer"] = struct{}{}

	s.users.Store(req.Login, userInfo{
		password: req.Password,
		roles:    roles,
	})

	return &empty.Empty{}, nil
}

func (s *Server) AuthTokenHolder(ctx context.Context, req *replicator.AuthRequest) (*replicator.AuthResponse, error) {
	user, ok := s.users.Load(req.Login)

	if !ok {
		return nil, status.Error(codes.InvalidArgument, "token holder not registered")
	}

	ui := user.(userInfo)

	if ui.password != req.Password {
		return nil, status.Error(codes.InvalidArgument, "invalid password")
	}

	expire := time.Now().Add(time.Minute * 30)

	claims := loginCliams{
		req.Login,
		jwt.StandardClaims{
			ExpiresAt: jwt.NewTime(float64(expire.Unix())),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	jwtStore.Append(jwtstore.JWT{
		Token:       signedToken,
		ExpireDate:  expire,
		HolderLogin: req.Login,
	})

	return &replicator.AuthResponse{
		Jwt:        signedToken,
		ExpireDate: expire.Unix(),
	}, nil

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
	_, ok := tokens.Load(req.Offer.Token)
	if ok {
		return nil, status.Error(codes.InvalidArgument, "token with this name already exists")
	}

	tokens.Store(req.Offer.Token, token{
		price:     req.Offer.Price,
		validTime: req.Offer.ValidUntilSeconds,
	})

	return &emptypb.Empty{}, nil
}

func (s *Server) UpdateToken(ctx context.Context, req *replicator.UpdateTokenRequest) (*empty.Empty, error) {
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

func (s *Server) RevokeToken(ctx context.Context, req *replicator.RevokeTokenRequest) (*empty.Empty, error) {
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

func (s *Server) GetTokenList(ctx context.Context, req *replicator.GetTokenListRequest) (*replicator.GetTokenListResponse, error) {
	var tokenList []*replicator.TokenOffer
	var err error

	tokenList, err = s.allTokens()

	if err != nil {
		return nil, err
	}

	return &replicator.GetTokenListResponse{
		Tokens: tokenList,
		Total:  int32(len(tokenList)),
	}, nil

}

// TODO: Rework this method. Need geting all issuers and their tokens with wallet addresses
func (s *Server) allTokens() ([]*replicator.TokenOffer, error) {
	resultList := []*replicator.TokenOffer{}

	tokens.Range(func(key, value interface{}) bool {
		// TODO: Fill the resultList with token information
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
