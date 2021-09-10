package tokendb

import (
	"fmt"
	"os"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/channeldb"
	"github.com/pkt-cash/pktd/pktwallet/walletdb"
)

type TokenStrikeDB struct {
	client channeldb.DB
}

var ()

const (
	dbName = "channel.db"
)

// Function executed before every transaction
func before() {
	fmt.Println("Transaction")
}

func Connect(path string) (*TokenStrikeDB, error) {
	db, err := channeldb.Open(path)
	if err != nil {
		return nil, err.Native()
	}

	tsDB := &TokenStrikeDB{
		client: *db,
	}

	return tsDB, nil
}

func (t *TokenStrikeDB) Close() {
	t.client.Close()
}

func (t *TokenStrikeDB) Ping() er.R {
	return t.client.View(func(tx walletdb.ReadTx) er.R {
		return nil
	}, before)
}

// Clear all database structure with buckets and their content
func (t *TokenStrikeDB) Clear() er.R {
	dbFullPath := fmt.Sprintf("%v/%v", t.client.Path(), dbName)

	if err := os.Truncate(dbFullPath, 0); err != nil {
		return er.Errorf("Failed to truncate: %v", err)
	}
	t = nil
	return nil
}

func (t *TokenStrikeDB) Update(f func(tx walletdb.ReadWriteTx) er.R) er.R {
	return t.client.Update(f, before)
}

func (t *TokenStrikeDB) View(f func(tx walletdb.ReadTx) er.R) er.R {
	return t.client.View(f, before)
}
