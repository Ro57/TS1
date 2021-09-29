package main

import (
	"context"
	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/issuer"
	"github.com/urfave/cli"
)

var transferTokensCommand = cli.Command{
	Name:     "transfer-tokens",
	Category: "Tokens",
	Usage:    "transfer selected tokens by lock id using htlc secret",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "token",
			Usage: "selected token for transfer",
		},
		cli.Int64Flag{
			Name:  "count",
			Usage: "number of tokens to lock",
		},
		cli.StringFlag{
			Name:  "recipient",
			Usage: "wallet address of new owner of tokens",
		},
		cli.StringFlag{
			Name: "memo",
			Usage: "a description of the payment to attach along " +
				"with the invoice (default=\"\")",
		},
		cli.Int64Flag{
			Name:  "amt",
			Usage: "the amt of satoshis in this invoice",
		},
		cli.IntFlag{
			Name:  "proof_blocks",
			Usage: "lock expiration time in PKT blocks",
		},
	},
	Action: transferTokens,
}

func transferTokens(ctx *cli.Context) er.R {
	if !ctx.IsSet("recipient") {
		return er.New("recipient required")
	}

	if !ctx.IsSet("token") {
		return er.New("token required")
	}

	count := ctx.Int64("count")
	if count <= 0 {
		return er.New("count cannot be equal to or less than zero")
	}

	amt := ctx.Int64("amt")
	if amt <= 0 {
		return er.New("count cannot be equal to or less than zero")
	}

	proofBlocks := ctx.Int64("proofBlocks")
	if proofBlocks <= 0 {
		return er.New("proof_blocks cannot be equal to or less than zero")
	}

	req := &issuer.TransferTokensRequest{
		Token:      ctx.String("token"),
		Count:      count,
		Amt:        amt,
		Recipient:  ctx.String("recipient"),
		ProofCount: int32(proofBlocks),
		Memo:       ctx.String("memo"),
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
