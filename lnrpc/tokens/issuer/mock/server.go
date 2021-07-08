package mocks

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net"
	"sync"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lightninglabs/protobuf-hex-display/json"
	"github.com/pkg/errors"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/issuer"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	tokens sync.Map
)

type token struct {
	price       uint64
	validTime   int64
	issuerLogin string
}

type Server struct {
	// Nest unimplemented server implementation in order to satisfy server interface
	issuer.UnimplementedIssuerServiceServer

	Client replicator.ReplicatorClient
}

func RunServerServing(host string, replicationHost string, stopSig <-chan struct{}) {
	client, closeConn, err := connectReplicatorClient(context.TODO(), replicationHost)
	if err != nil {
		panic(err)
	}

	var (
		child = &Server{
			Client: client,
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
		<-stopSig
		root.Stop()
		closeConn()
	}()
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
	bytes, err := json.Marshal(req)
	if err != nil {
		return nil, errors.WithMessage(err, "marshalling request")
	}

	hash := sha256.Sum256(bytes)

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
