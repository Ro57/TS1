package main

import (
	"context"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/replicator"
	"github.com/urfave/cli"
)

var verifyTokenPurchaseSignatureCommand = cli.Command{
	Name:        "verifytokenpurchasesignature",
	Category:    "Tokens",
	Usage:       "Verifies issuer's token purchase signature",
	Description: "Verifies issuer's token purchase signature. Successfull call on this method grants, that the token purchase can be registered by Replication Server",
	Flags:       verifyTokenPurchaseSignatureFlags,
	Action:      actionDecorator(verifyTokenPurchaseSignature),
}

const (
	flagIssuerTokenPurchaseSignature = "issuer-signature"
)

var verifyTokenPurchaseSignatureFlags = append(
	[]cli.Flag{
		cli.StringFlag{
			Name:  flagIssuerTokenPurchaseSignature,
			Usage: "signature maden by responsible issuer, presenting issuer's good will to follow purchase terms",
		},
	},
	signTokenPurchaseFlags...,
)

func verifyTokenPurchaseSignature(ctx *cli.Context) er.R {
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

	req := &replicator.VerifyTokenPurchaseRequest{
		Purchase: &replicator.TokenPurchase{
			Offer:           offer,
			IssuerSignature: issuerSignature,
		},
	}
	_, err := client.VerifyTokenPurchase(context.TODO(), req)
	if err != nil {
		return er.Errorf("requesting token purchase signature verification: %s", err)
	}

	println("OK")

	return nil
}
