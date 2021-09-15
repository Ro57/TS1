package main

import (
	"context"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"github.com/urfave/cli"
)

var getHeadersCommand = cli.Command{
	Name:        "get-headers",
	Category:    "Tokens",
	Usage:       "List information about available token set per token id.",
	Description: `List information about available token offers per token id.`,

	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "tokenID",
			Usage: "Returned token info of would belong token id",
		},
		cli.StringFlag{
			Name:  "hash",
			Usage: "",
		},
	},
	Action: getHeaders,
}

func getHeaders(ctx *cli.Context) er.R {
	client, cleanUp := getClient(ctx)
	defer cleanUp()

	var ( // Default request parameters - no pagination
		tokenID string
		hash    string
	)

	// Acquire passed values, that are not zero
	tokenID = ctx.String("tokenID")
	if tokenID == "" {
		return er.New("token id cannot is empty")
	}

	hash = ctx.String("hash")
	if tokenID == "" {
		return er.New("hash cannot is empty")
	}

	// Request offers
	req := &replicator.GetHeadersRequest{
		TokenId: tokenID,
		Hash:    hash,
	}
	resp, err := client.GetHeaders(context.TODO(), req)
	if err != nil {
		return er.E(err)
	}

	printRespJSON(resp)

	return nil
}
