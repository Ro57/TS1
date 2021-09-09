package hashhelper

import (
	"fmt"
	"strings"
)

type Address string
type Transaction string
type Block string
type Lock string

func NewAddress(hash string) (*Address, error) {
	a := Address(hash)
	if !a.Validate() {
		return nil, fmt.Errorf("hash %v not a address", hash)
	}

	return &a, nil
}

func NewBlock(hash string) (*Block, error) {
	b := Block(hash)
	if !b.Validate() {
		return nil, fmt.Errorf("hash %v not a block", hash)
	}

	return &b, nil
}

func NewTransaction(hash string) (*Transaction, error) {
	t := Transaction(hash)
	if !t.Validate() {
		return nil, fmt.Errorf("hash %v not a transaction", hash)
	}

	return &t, nil
}

func NewLock(hash string) (*Lock, error) {
	l := Lock(hash)
	if !l.Validate() {
		return nil, fmt.Errorf("hash %v not a lock", hash)
	}

	return &l, nil
}

func (a *Address) Validate() bool {
	return strings.HasPrefix(string(*a), "pkt1")
}

func (t *Transaction) Validate() bool {
	return strings.HasPrefix(string(*t), "tx")
}

func (b *Block) Validate() bool {
	return strings.HasPrefix(string(*b), "block")
}

func (l *Lock) Validate() bool {
	return strings.HasPrefix(string(*l), "lock")
}
