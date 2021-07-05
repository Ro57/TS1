package main

import (
	"context"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"github.com/urfave/cli"
)

var registerTokenPurchaseCommand = cli.Command{
	Name:        "registertokenpurchase",
	Category:    "Tokens",
	Usage:       "Registers initialized token purchase at Replication server",
	Description: "Registers initialized token purchase at Replication server to be protected by",
	Flags:       registerTokenPurchaseFlags,
	Action:      actionDecorator(registerTokenPurchase),
}

const (
	flagInitialTxHash = "initial-tx-hash"
)

var registerTokenPurchaseFlags = append(
	[]cli.Flag{
		cli.StringFlag{
			Name:  flagInitialTxHash,
			Usage: "channel opening transaction hash",
		},
	},
	verifyTokenPurchaseSignatureFlags...,
)

func registerTokenPurchase(ctx *cli.Context) er.R {
	client, cleanUp := getClient(ctx)
	defer cleanUp()

	offer, _err := extractTokenOffer(ctx)
	if _err != nil {
		return _err
	}

	issuerSignature := ctx.String(flagIssuerTokenPurchaseSignature)
	if issuerSignature == "" {
		return er.Errorf("empty %q argument provided", flagIssuerTokenPurchaseSignature)
	}

	initialTxHash := ctx.String(flagInitialTxHash)
	if initialTxHash == "" {
		return er.Errorf("empty %q argument provided", flagInitialTxHash)
	}

	req := &replicator.RegisterTokenPurchaseRequest{
		Purchase: &replicator.TokenPurchase{
			Offer:           offer,
			IssuerSignature: issuerSignature,
			InitialTxHash:   initialTxHash,
		},
	}
	_, err := client.RegisterTokenPurchase(context.TODO(), req)
	if err != nil {
		return er.Errorf("requesting token purchase registration: %s", err)
	}

	println("OK")

	return nil
}
