package hashhelper

import (
	"fmt"
	"strings"
)

type address string
type transaction string
type block string
type lock string

func NewAddress(hash string) (*address, error) {
	a := address(hash)
	if !a.validate() {
		return nil, fmt.Errorf("hash %v not a address", hash)
	}

	return &a, nil
}

func NewBlock(hash string) (*block, error) {
	b := block(hash)
	if !b.validate() {
		return nil, fmt.Errorf("hash %v not a block", hash)
	}

	return &b, nil
}

func NewTransaction(hash string) (*transaction, error) {
	t := transaction(hash)
	if !t.validate() {
		return nil, fmt.Errorf("hash %v not a transaction", hash)
	}

	return &t, nil
}

func NewLock(hash string) (*lock, error) {
	l := lock(hash)
	if !l.validate() {
		return nil, fmt.Errorf("hash %v not a lock", hash)
	}

	return &l, nil
}

func (a *address) validate() bool {
	return strings.HasPrefix(string(*a), "pkt1")
}

func (t *transaction) validate() bool {
	return strings.HasPrefix(string(*t), "tx")
}

func (b *block) validate() bool {
	return strings.HasPrefix(string(*b), "block")
}

func (l *lock) validate() bool {
	return strings.HasPrefix(string(*l), "lock")
}
