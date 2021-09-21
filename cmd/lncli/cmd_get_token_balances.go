package main

import (
	"context"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"github.com/urfave/cli"
)

var getTokenBalancesCommand = cli.Command{
	Name:     "gettokenbalances",
	Category: "Tokens",
	Usage:    "List information about current token balances.",
	Description: `List information about current token balances. 

	There is an opportunity to list token balances in a pagination-like manner. A such behaviour 
can be achieved by providing additional flags to the command.
DEPRECATED `,
	Flags: []cli.Flag{
		cli.UintFlag{
			Name:  "limit",
			Usage: "(optional) If a value provided, returned token balances number would be limited to the specified value",
		},
		cli.UintFlag{
			Name:  "offset",
			Usage: "(optional) If a value provided, returned token balances would be at the specified offset \"height\"",
		},
	},
	Action: getTokenBalances,
}

func getTokenBalances(ctx *cli.Context) er.R {
	client, cleanUp := getClient(ctx)
	defer cleanUp()

	var ( // Default request parameters - no pagination
		limit  uint64
		offset uint64
	)

	// Acquire passed values, that are not zero
	if v := ctx.Uint64("limit"); v != 0 {
		limit = v
	}
	if v := ctx.Uint64("offset"); v != 0 {
		offset = v
	}

	// Request token balances
	req := &lnrpc.GetTokenBalancesRequest{
		Params: &replicator.Pagination{
			Limit:  limit,
			Offset: offset,
		},
	}

	resp, err := client.GetTokenBalances(context.TODO(), req)
	if err != nil {
		return er.E(err)
	}
	printRespJSON(resp)

	return nil
}
