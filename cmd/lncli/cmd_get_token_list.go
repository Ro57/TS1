package main

import (
	"context"
	"fmt"
	"log"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/channeldb"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"github.com/pkt-cash/pktd/pktwallet/walletdb"
	"github.com/urfave/cli"
)

var getTokenListCommand = cli.Command{
	Name:     "issuedtokens",
	Category: "Tokens",
	Usage:    "List information about available token set per issuer.",
	Description: `List information about available token offers per issuer. 

	"Available offers" means such offers, that officially registered on the off-chain ecosystem and 
all the related deals would be tracked and protected by an overseer. 

	There is an opportunity to list available offers in a pagination-like manner. A such behaviour 
can be achieved by providing additional flags to the command.`,

	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "issuer-id",
			Usage: "(optional) If a value provided, returned offers would belong to the specified token issuer",
		},
		cli.UintFlag{
			Name:  "limit",
			Usage: "(optional) If a value provided, returned offers number would be limited to the specified value",
		},
		cli.UintFlag{
			Name:  "offset",
			Usage: "(optional) If a value provided, returned offers would be at the specified offset \"height\"",
		},
	},
	Action: getTokenList,
}

func getTokenList(ctx *cli.Context) er.R {
	client, cleanUp := getClient(ctx)
	defer cleanUp()

	var ( // Default request parameters - no pagination
		limit    uint64
		offset   uint64
		issuerID string
	)

	// Acquire passed values, that are not zero
	if v := ctx.Uint64("limit"); v != 0 {
		limit = v
	}
	if v := ctx.Uint64("offset"); v != 0 {
		offset = v
	}
	if v := ctx.String("issuer-id"); v != "" {
		issuerID = v
	}

	// Request offers
	req := &replicator.GetTokenListRequest{
		IssuerId: issuerID,
		Params: &replicator.Pagination{
			Limit:  limit,
			Offset: offset,
		},
	}
	resp, err := client.GetTokenList(context.TODO(), req)
	if err != nil {
		return er.E(err)
	}

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

	printRespJSON(resp)

	return nil
}
