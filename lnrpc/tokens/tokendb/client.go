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

// TODO: Change to config data

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
	}, func() {})
}

// Clear all database structure with buckets and their content
func Clear() er.R {
	dbFullPath := fmt.Sprintf("%v/%v", client.Path(), dbName)

	if err := os.Truncate(dbFullPath, 0); err != nil {
		return er.Errorf("Failed to truncate: %v", err)
	}

	return nil
}

// DB example
// func ExampleDB() {

// 	client, errr := channeldb.Open(home + path + name)
// 	if errr != nil {
// 		log.Fatal(errr.Native().Error())
// 	}
// 	defer db.Close()

// 	db.Update(func(tx walletdb.ReadWriteTx) er.R {
// 		b, err := tx.CreateTopLevelBucket([]byte("Business"))
// 		if err != nil {
// 			return er.Errorf("create bucket: %s", err)
// 		}

// 		_, err = b.CreateBucket([]byte("Employee"))
// 		if err != nil {
// 			return er.Errorf("create bucket: %s", err)
// 		}

// 		return nil

// 	}, func() {
// 		fmt.Println("create busket structure")
// 	})

// 	db.Update(func(tx walletdb.ReadWriteTx) er.R {
// 		b := tx.ReadWriteBucket([]byte("Business"))
// 		emp := b.NestedReadWriteBucket([]byte("Employee"))

// 		err := b.Put([]byte("Emp1"), []byte("Number 100"))

// 		if err != nil {
// 			return err
// 		}

// 		err = emp.Put([]byte("Time"), []byte("180h"))

// 		return err
// 	}, func() {
// 		fmt.Println("Put Emp1")
// 	})

// 	db.View(func(tx walletdb.ReadTx) er.R {
// 		b := tx.ReadBucket([]byte("Business"))
// 		emp := b.NestedReadBucket([]byte("Employee"))

// 		empNum := b.Get([]byte("Emp1"))

// 		empTime := emp.Get([]byte("Time"))
// 		fmt.Printf("Hello %s your working time %s", empNum, empTime)

// 		return nil
// 	}, func() {
// 	})
// }
