package main

import (
	"context"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"github.com/urfave/cli"
)

var registerTokenSellCommand = cli.Command{
	Name:        "registertokensell",
	Category:    "Tokens",
	Usage:       "Registers initialized token sell at Replication server",
	Description: "Registers initialized token sell at Replication server to be protected by",
	Flags:       registerTokenSellFlags,
	Action:      actionDecorator(registerTokensell),
}

var registerTokenSellFlags = append(
	[]cli.Flag{
		cli.StringFlag{
			Name:  flagInitialTxHash,
			Usage: "channel opening transaction hash",
		},
	},
	verifyTokenSellSignatureFlags...,
)

func registerTokensell(ctx *cli.Context) er.R {
	client, cleanUp := getClient(ctx)
	defer cleanUp()

	offer, _err := extractTokenSell(ctx)
	if _err != nil {
		return _err
	}

	issuerSignature, err := parseRequiredString(ctx, flagIssuerTokenSellSignature)
	if err != nil {
		return er.E(err)
	}

	initialTxHash, err := parseRequiredString(ctx, flagInitialTxHash)
	if err != nil {
		return er.E(err)
	}

	req := &replicator.RegisterTokenSellRequest{
		Sell: &replicator.TokenPurchase{
			Offer:           offer,
			IssuerSignature: issuerSignature,
			InitialTxHash:   initialTxHash,
		},
	}

	_, err = client.RegisterTokenSell(context.TODO(), req)
	if err != nil {
		return er.Errorf("requesting token sell registration: %s", err)
	}

	println("OK")

	return nil
}
