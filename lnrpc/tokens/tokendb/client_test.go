package tokendb

import (
	"fmt"
	"testing"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/pktwallet/walletdb"
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

func TestUpdate(t *testing.T) {
	Connect(path + name)
	defer Close()

	defer func() {
		err := Clear()
		if err != nil {
			t.Fatal("Clear failed", err.Native())
		}
	}()

	testDB := GetClient()
	err := testDB.Update(func(tx walletdb.ReadWriteTx) er.R {
		b, err := tx.CreateTopLevelBucket([]byte("Business"))

		if err != nil {
			return er.Errorf("create bucket: %s", err)
		}

		_, err = b.CreateBucket([]byte("Employee"))
		if err != nil {
			return er.Errorf("create bucket: %s", err)
		}

		return nil

	}, func() {
		fmt.Println("create busket structure")
	})

	if err != nil {
		t.Fatal("Create busket structures failed: ", err)
	}

}
