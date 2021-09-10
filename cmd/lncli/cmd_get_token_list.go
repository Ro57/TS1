package main

import (
	"context"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"github.com/urfave/cli"
)

// TODO: This command now breaked, need research server logic and repair it.
var getTokenListCommand = cli.Command{
	Name:     "get-token-list",
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

	printRespJSON(resp)

	return nil
}

var getTokenCommand = cli.Command{
	Name:        "get-token",
	Category:    "Tokens",
	Usage:       "List information about available token set per token id.",
	Description: `List information about available token offers per token id.`,

	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "tokenID",
			Usage: "Returned token info of would belong token id",
		},
	},
	Action: getToken,
}

func getToken(ctx *cli.Context) er.R {
	client, cleanUp := getClient(ctx)
	defer cleanUp()

	var ( // Default request parameters - no pagination
		tokenID string
	)

	// Acquire passed values, that are not zero
	if v := ctx.String("tokenID"); v != "" {
		tokenID = v
	}

	// Request offers
	req := &replicator.GetTokenRequest{
		TokenId: tokenID,
	}
	resp, err := client.GetToken(context.TODO(), req)
	if err != nil {
		return er.E(err)
	}

	printRespJSON(resp)

	return nil
}
