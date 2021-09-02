package db

import (
	"fmt"
	"log"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/channeldb"
	"github.com/pkt-cash/pktd/pktwallet/walletdb"
)

// DB example
func exampleDB() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, errr := channeldb.Open("tokens")
	if errr != nil {
		log.Fatal(errr.Native().Error())
	}
	defer db.Close()

	db.Update(func(tx walletdb.ReadWriteTx) er.R {
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

	db.Update(func(tx walletdb.ReadWriteTx) er.R {
		b := tx.ReadWriteBucket([]byte("Business"))
		emp := b.NestedReadWriteBucket([]byte("Employee"))

		err := b.Put([]byte("Emp1"), []byte("Number 100"))

		if err != nil {
			return err
		}

		err = emp.Put([]byte("Time"), []byte("180h"))

		return err
	}, func() {
		fmt.Println("Put Emp1")
	})

	db.View(func(tx walletdb.ReadTx) er.R {
		b := tx.ReadBucket([]byte("Business"))
		emp := b.NestedReadBucket([]byte("Employee"))

		empNum := b.Get([]byte("Emp1"))

		empTime := emp.Get([]byte("Time"))
		fmt.Printf("Hello %s your working time %s", empNum, empTime)

		return nil
	}, func() {
	})
}
