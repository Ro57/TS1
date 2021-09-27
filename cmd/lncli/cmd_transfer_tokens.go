package main

import (
	"context"
	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc"
	"github.com/urfave/cli"
)

var transferTokensCommand = cli.Command{
	Name:     "transfer-tokens",
	Category: "Tokens",
	Usage:    "transfer selected tokens by lock id using htlc secret",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "lockID",
			Usage: "id of selected lock of tokens",
		},
		cli.StringFlag{
			Name:  "secret",
			Usage: "htlc secret",
		},
		cli.StringFlag{
			Name:  "token",
			Usage: "selected token for transfer",
		},
	},
	Action: transferTokens,
}

func transferTokens(ctx *cli.Context) er.R {
	if !ctx.IsSet("lockID") {
		return er.New("lockID required")
	}

	if !ctx.IsSet("secret") {
		return er.New("htlc secret required")
	}

	if !ctx.IsSet("token") {
		return er.New("token required")
	}

	req := &lnrpc.TokenTransfersRequest{
		LockID:     ctx.String("lockID"),
		HtlcSecret: ctx.String("secret"),
		Token:      ctx.String("token"),
	}

	client, cleanUp := getClient(ctx)
	defer cleanUp()

	resp, err := client.TransferTokens(context.TODO(), req)
	if err != nil {
		return er.E(err)
	}
	printRespJSON(resp)

	return nil
}
