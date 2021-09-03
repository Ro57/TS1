package tokendb

import (
	"crypto/sha256"
	"encoding/json"
	"testing"
	"time"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/DB"
	"github.com/pkt-cash/pktd/pktwallet/walletdb"
	"google.golang.org/protobuf/types/known/anypb"
)

const (
	path = "/.lnd/data/chain/pkt"
	name = "/tokens"
)

func TestCreateDB(t *testing.T) {
	Connect(path + name)
	defer Close()
	err := Ping()

	if err != nil {
		t.Fatal("db not connected: ", err)
	}
}

func TestPing(t *testing.T) {
	Connect(path + name)
	Close()
	err := Ping()

	if err == nil {
		t.Fatal("ping after close connection")
	}
}

func TestEmployeeUpdateView(t *testing.T) {
	Connect(path + name)
	defer Close()

	defer func() {
		err := Clear()
		if err != nil {
			t.Fatal("Clear failed", err.Native())
		}
	}()

	err := Update(func(tx walletdb.ReadWriteTx) er.R {
		_, err := tx.CreateTopLevelBucket([]byte("Business"))

		if err != nil {
			return er.Errorf("create bucket: %s", err)
		}

		return nil
	})
	if err != nil {
		t.Fatal("Create top level busket structures failed: ", err)
	}

	err = Update(func(tx walletdb.ReadWriteTx) er.R {
		b := tx.ReadWriteBucket([]byte("Business"))

		_, err = b.CreateBucket([]byte("Employee"))
		if err != nil {
			return er.Errorf("create bucket: %s", err)
		}

		return nil

	})
	if err != nil {
		t.Fatal("Create busket structures failed: ", err)
	}

	err = Update(func(tx walletdb.ReadWriteTx) er.R {
		b := tx.ReadWriteBucket([]byte("Business"))
		emp := b.NestedReadWriteBucket([]byte("Employee"))

		err := b.Put([]byte("Emp1"), []byte("Number 100"))

		if err != nil {
			return err
		}

		err = emp.Put([]byte("Time"), []byte("180h"))

		return err
	})
	if err != nil {
		t.Fatal("Put information into bucket: ", err)
	}

	err = View(func(tx walletdb.ReadTx) er.R {
		b := tx.ReadBucket([]byte("Business"))
		emp := b.NestedReadBucket([]byte("Employee"))

		empNum := b.Get([]byte("Emp1"))
		if string(empNum) != "Number 100" {
			t.Fatalf("want string 'Number 100' but get '%s'", string(empNum))
		}

		empTime := emp.Get([]byte("Time"))
		if string(empTime) != "180h" {
			t.Fatalf("want string '180h' but get '%s'", string(empTime))
		}

		return nil
	})
	if err != nil {
		t.Fatal("Get information from bucket: ", err)
	}
}

func TestTockenBlock(t *testing.T) {
	var wantExpBlockNumber int32 = 2000
	wantCreateTime := time.Now()
	tokenName := "smt"
	wantToken := DB.Token{
		Count:      200,
		Expiration: wantExpBlockNumber,
		Creation:   wantCreateTime.UnixNano(),
		Url:        "",
	}

	wantBlock := DB.Block{
		// TODO: Change on real justification structure after declaration
		Justification:  &anypb.Any{Value: []byte("LOCK_TOKEN")},
		Signature:      "SomeSig",
		State:          "SomeHash",
		AvailableCount: 200,
		Locks: []*DB.Lock{
			{
				Id:         "SomeHash",
				Count:      1,
				Owner:      "pkt1otherwallet",
				Htlc:       "someHTLCHash",
				ProofCount: 2000,
			},
		},
		Owners: []*DB.Owner{
			{
				HolderWallet: "pkt1somewallet",
				Count:        1,
			},
		},
	}

	var lastBlock [sha256.Size]byte

	Connect(path + name)
	defer Close()

	defer func() {
		err := Clear()
		if err != nil {
			t.Fatal("Clear failed", err.Native())
		}
	}()

	err := Update(func(tx walletdb.ReadWriteTx) er.R {
		b, err := tx.CreateTopLevelBucket([]byte(tokenName))
		if err != nil {
			t.Fatalf("create top level bucket: %s", err)
		}

		chain, err := b.CreateBucket([]byte("chain"))
		if err != nil {
			t.Fatalf("create chain bucket: %s", err)
		}

		tokenByte, nativeErr := json.Marshal(wantToken)
		if nativeErr != nil {
			t.Fatalf("(update) marshal token structure: %s", nativeErr)
		}

		err = b.Put([]byte("Info"), tokenByte)
		if err != nil {
			t.Fatalf("put token %s", err.Native())
		}

		blockByte, nativeErr := json.Marshal(wantBlock)
		if nativeErr != nil {
			t.Fatalf("(update) marshal block structure: %s", nativeErr)
		}

		lastBlock = sha256.Sum256(blockByte)

		err = chain.Put(lastBlock[:], blockByte)
		if err != nil {
			t.Fatalf("put block %s", err.Native())
		}

		return nil
	})
	if err != nil {
		t.Fatal("generate DB structure: ", err)
	}

	err = View(func(tx walletdb.ReadTx) er.R {
		b := tx.ReadBucket([]byte(tokenName))
		if b == nil {
			t.Fatalf("bucket %s not found", tokenName)
		}

		chain := b.NestedReadBucket([]byte("chain"))
		if chain == nil {
			t.Fatal("bucket chain not found")
		}

		marshalToken, nativeErr := json.Marshal(wantToken)
		if nativeErr != nil {
			t.Fatalf("(view) marshal token structure: %s", nativeErr)
		}

		tokenByte := b.Get([]byte("Info"))
		if tokenByte == nil {
			t.Fatalf("token info with name %s not found", tokenByte)
		}
		if string(marshalToken) != string(tokenByte) {
			t.Fatalf("want token structure %s but get %s", marshalToken, tokenByte)
		}

		marshalBlock, nativeErr := json.Marshal(wantBlock)
		if nativeErr != nil {
			t.Fatalf("(view) marshal token structure: %s", nativeErr)
		}

		blockByte := chain.Get(lastBlock[:])
		if blockByte == nil {
			t.Fatalf("block with hash %s not found", string(lastBlock[:]))
		}
		if string(marshalBlock) != string(blockByte) {
			t.Fatalf("want block structure %s but get %s", marshalBlock, blockByte)
		}

		return nil
	})
	if err != nil {
		t.Fatal("read data from DB: ", err)
	}

}
