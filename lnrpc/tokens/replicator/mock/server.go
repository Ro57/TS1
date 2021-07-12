package mock

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/jwtstore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// token holders with login password
var (
	users      sync.Map
	jwtStore   *jwtstore.Store
	signingKey = []byte("SUPER_SECRET")
)

type userInfo struct {
	password string
	// TODO: implement role system
	// ? method for add new role to user
	roles map[string]struct{}
}

type ReplicatorEvents struct {
	StopSig chan struct{}
}

type Server struct {
	// Nest unimplemented server implementation in order to satisfy server interface
	replicator.UnimplementedReplicatorServer

	holders        TokenHoldersStoreAPI
	holderBalances TokenHolderBalancesStoreAPI

	events ReplicatorEvents
}

type loginCliams struct {
	login string
	jwt.StandardClaims
}

func RunServerServing(host string, events ReplicatorEvents) {
	var (
		child = &Server{
			holders:        NewTokenHoldersStore(),
			holderBalances: NewTokenHolderBalancesStore(),
			events:         events,
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
	const (
		tokensNum = 100
	)

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
		return nil, status.Error(codes.ResourceExhausted, fmt.Sprintf("session expired"))
	}

	balances := make([]*replicator.TokenBalance, 0, tokensNum)

	// Fill mocked token balances. At this time balances are not owned by a specific holder
	for i := tokensNum; i > 0; i-- {
		balance := &replicator.TokenBalance{
			Token:     fmt.Sprintf("token_%d", i),
			Available: uint64(i*2 + 1),
			Frozen:    uint64(i*3 + 1),
		}
		balances = append(balances, balance)
	}

	resp := &replicator.GetTokenBalancesResponse{
		Balances: balances,
		Total:    tokensNum,
	}

	// Apply pagination
	if req.Params.Offset > 0 {
		if int(req.Params.Offset) <= len(resp.Balances)-1 {
			resp.Balances = resp.Balances[req.Params.Offset:]
		} else {
			resp.Balances = nil
		}
	}

	if req.Params.Limit > 0 {
		if int(req.Params.Limit) <= len(resp.Balances)-1 {
			resp.Balances = resp.Balances[:req.Params.Limit]
		}
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

	return &empty.Empty{}, nil
}

func (s *Server) RegisterTokenPurchase(ctx context.Context, req *replicator.RegisterTokenPurchaseRequest) (*empty.Empty, error) {
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

	return &empty.Empty{}, nil
}

func (s *Server) RegisterTokenHolder(ctx context.Context, req *replicator.RegisterRequest) (*empty.Empty, error) {
	_, ok := users.Load(req.Login)
	if ok {
		return nil, status.Error(codes.InvalidArgument, "user with this login already exists")
	}

	roles := make(map[string]struct{})
	roles["holder"] = struct{}{}

	users.Store(req.Login, userInfo{
		password: req.Password,
		roles:    roles,
	})

	return &empty.Empty{}, nil
}

func (s *Server) RegisterTokenIssuer(ctx context.Context, req *replicator.RegisterRequest) (*empty.Empty, error) {
	_, ok := users.Load(req.Login)
	if ok {
		return nil, status.Error(codes.InvalidArgument, "user with this login already exists")
	}
	roles := make(map[string]struct{})
	roles["issuer"] = struct{}{}

	users.Store(req.Login, userInfo{
		password: req.Password,
		roles:    roles,
	})

	return &empty.Empty{}, nil
}

func (s *Server) AuthTokenHolder(ctx context.Context, req *replicator.AuthRequest) (*replicator.AuthResponse, error) {
	user, ok := users.Load(req.Login)

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
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjM5MjI1OTksImlzcyI6IjExMSJ9.eaSaGGyxzOLsGaAcQCQgKUwbwKvh9xy35lef1xF89u8

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

	user, ok := users.Load(req.Login)
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

type TokenHolder struct {
	Login    TokenHolderLogin
	Password string
}

type TokenHolderBalancesStoreAPI interface {
	Set(TokenHolderLogin, TokenHolderBalances) error
	Get(TokenHolderLogin) TokenHolderBalances
}
type TokenHolderBalancesStore struct {
	holders map[TokenHolderLogin]TokenHolderBalances
	mu      sync.RWMutex
}

var _ TokenHolderBalancesStoreAPI = (*TokenHolderBalancesStore)(nil)

func NewTokenHolderBalancesStore() *TokenHolderBalancesStore {
	return &TokenHolderBalancesStore{
		holders: make(map[TokenHolderLogin]TokenHolderBalances),
	}
}

func (s *TokenHolderBalancesStore) Set(login TokenHolderLogin, balances TokenHolderBalances) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.holders[login] = balances

	return nil
}

func (s *TokenHolderBalancesStore) Get(login TokenHolderLogin) TokenHolderBalances {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.holders[login]
}

type TokenHolderBalances []replicator.TokenBalance
