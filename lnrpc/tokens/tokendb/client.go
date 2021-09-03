package tokendb

import (
	"fmt"
	"os"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/channeldb"
	"github.com/pkt-cash/pktd/pktwallet/walletdb"
)

var (
	client channeldb.DB
)

const (
	dbName = "channel.db"
)

// Function executed before every transaction
func before() {
	fmt.Println("Transaction")
}
func GetClient() channeldb.DB {
	return client
}

func Connect(path string) error {
	home, _ := os.UserHomeDir()
	db, err := channeldb.Open(home + path)
	if err != nil {
		return err.Native()
	}

	client = *db

	return nil
}

func Close() {
	client.Close()
}

func Ping() er.R {
	return client.View(func(tx walletdb.ReadTx) er.R {
		return nil
	}, before)
}

// Clear all database structure with buckets and their content
func Clear() er.R {
	dbFullPath := fmt.Sprintf("%v/%v", client.Path(), dbName)

	if err := os.Truncate(dbFullPath, 0); err != nil {
		return er.Errorf("Failed to truncate: %v", err)
	}

	return nil
}

func Update(f func(tx walletdb.ReadWriteTx) er.R) er.R {
	return client.Update(f, before)
}

func View(f func(tx walletdb.ReadTx) er.R) er.R {
	return client.View(f, before)
}
