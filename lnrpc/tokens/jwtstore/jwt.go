package jwtstore

// JWTStore in memory storage of all authorized token with user id
// and payload

import (
	"errors"
	"sync"
	"time"
)

// JWT is structured datas of token information
// returned from replicator serive
type JWT struct {
	Token       string
	ExpireDate  time.Time
	HolderLogin string
}

// Store contains all token information and implements methods
// of sync access to this data
type Store struct {
	mu     sync.RWMutex
	tokens []JWT
}

var errLoginNotFound = errors.New(`
JWT with given holder login not found.
Perhaps this session was deleted when the server was restarted and you need a new sesssion`)

var errTokenNotFound = errors.New(`
JWT with given token not found.
Perhaps this session was deleted when the server was restarted and you need a new sesssion`)

func New(tokens []JWT) *Store {
	store := new(Store)
	store.tokens = tokens

	return store
}

func (s *Store) Remove(login string) error {
	index, err := s.position(login)
	if err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.tokens = append(s.tokens[:index], s.tokens[index+1:]...)
	return nil
}

func (s *Store) GetByLogin(login string) (JWT, error) {

	pos, err := s.position(login)
	if err != nil {
		return JWT{}, err
	}

	return s.tokens[pos], nil
}

func (s *Store) GetByToken(token string) (JWT, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, t := range s.tokens {
		if t.Token == token {
			return t, nil
		}
	}

	return JWT{}, errTokenNotFound
}

func (s *Store) Append(token JWT) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.tokens = append(s.tokens, token)
}

func (s *Store) position(login string) (int, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for i, t := range s.tokens {
		if t.HolderLogin == login {
			return i, nil
		}
	}

	return 0, errLoginNotFound
}
