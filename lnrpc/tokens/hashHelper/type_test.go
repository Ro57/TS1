package hashhelper

import "testing"

func TestAddress(t *testing.T) {
	a, err := NewAddress("pkt1qjucsc48vg6njy9djp4utwyzatgk6qdzeu9unmg")
	if err != nil {
		t.Fatalf("on create address %v", err)
	}

	isAddress := a.validate()
	if !isAddress {
		t.Fatalf("hash %v created as a address", a)
	}

}

func TestBlock(t *testing.T) {
	b, err := NewBlock("blocksome")
	if err != nil {
		t.Fatalf("on create block %v", err)
	}

	isBlock := b.validate()
	if !isBlock {
		t.Fatalf("hash %v created as a block", b)
	}
}

func TestTransaction(t *testing.T) {
	tx, err := NewTransaction("txsome")
	if err != nil {
		t.Fatalf("on create transaction %v", err)
	}

	isTransaction := tx.validate()
	if !isTransaction {
		t.Fatalf("hash %v created as a transaction", tx)
	}
}

func TestLock(t *testing.T) {
	l, err := NewLock("locksome")
	if err != nil {
		t.Fatalf("on create lock %v", err)
	}

	isLock := l.validate()
	if !isLock {
		t.Fatalf("hash %v created as a lock", l)
	}
}
