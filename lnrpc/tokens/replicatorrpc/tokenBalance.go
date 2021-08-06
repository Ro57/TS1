package replicatorrpc

import (
	"github.com/pkg/errors"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"github.com/pkt-cash/pktd/pktlog/log"
)

type TokenHolderBalancesStoreAPI interface {
	Get(TokenName) (*replicator.TokenBalance, error)
	Upadte(TokenName, *replicator.TokenBalance) error
	Append(TokenHolderBalances) error
	AppendOrUpdate(TokenHolderBalances)
	IsContain(TokenName) bool
}

type TokenHolderBalancesStore struct {
	store TokenHolderBalances
}

var _ TokenHolderBalancesStoreAPI = (*TokenHolderBalancesStore)(nil)

func NewTokenHolderBalancesStore(balances TokenHolderBalances) *TokenHolderBalancesStore {
	return &TokenHolderBalancesStore{
		store: balances,
	}
}

func (s *TokenHolderBalancesStore) IsContain(token TokenName) bool {
	for _, v := range s.store {
		if v.Token == token {
			return true
		}
	}

	return false
}

func (s *TokenHolderBalancesStore) Get(token TokenName) (*replicator.TokenBalance, error) {
	for _, v := range s.store {
		if v.Token == token {
			return v, nil
		}
	}

	return nil, errors.Errorf("Token with %v name not found in store", token)
}

func (s *TokenHolderBalancesStore) Append(balances TokenHolderBalances) error {
	for _, v := range balances {
		if s.IsContain(v.Token) {
			return errors.Errorf("Failed to append token. Token with %v name is already in store.", v.Token)
		}
	}

	s.store = append(s.store, balances...)
	return nil
}

func (s *TokenHolderBalancesStore) AppendOrUpdate(balances TokenHolderBalances) {
	for _, v := range balances {
		s.appendOrUpdate(v)
	}
}

func (s *TokenHolderBalancesStore) appendOrUpdate(balance *replicator.TokenBalance) {

	if s.IsContain(balance.Token) {
		log.Info(balance.Available)

		token, _ := s.Get(balance.Token)
		s.Upadte(token.Token, &replicator.TokenBalance{
			Token:     token.Token,
			Available: token.Available + balance.Available,
			Frozen:    token.Frozen + balance.Frozen,
		})
		return
	}

	s.Append([]*replicator.TokenBalance{balance})
}
func (s *TokenHolderBalancesStore) Upadte(token TokenName, balance *replicator.TokenBalance) error {
	for i, v := range s.store {
		if v.Token == token {
			s.store[i] = balance
		}
	}

	return errors.Errorf("Token with %v name not found in store", token)
}

type TokenHolderBalances = []*replicator.TokenBalance
