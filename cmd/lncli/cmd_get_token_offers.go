package main

import (
	"context"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"github.com/urfave/cli"
)

var getTokenOffersCommand = cli.Command{
	Name:     "gettokenoffers",
	Category: "Tokens",
	Usage:    "List information about available token offers per issuer.",
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
	Action: getTokenOffers,
}

func getTokenOffers(ctx *cli.Context) er.R {
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
	req := &replicator.GetTokenOffersRequest{
		IssuerId: issuerID,
		Params: &replicator.Pagination{
			Limit:  limit,
			Offset: offset,
		},
	}
	resp, err := client.GetTokenOffers(context.TODO(), req)
	if err != nil {
		return er.E(err)
	}
	printRespJSON(resp)

	return nil
}
