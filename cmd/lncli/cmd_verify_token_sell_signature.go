package main

import (
	"context"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/replicator"
	"github.com/urfave/cli"
)

var verifyTokenSellSignatureCommand = cli.Command{
	Name:        "verifytokensellsignature",
	Category:    "Tokens",
	Usage:       "Verifies issuer's token sell signature",
	Description: "Verifies issuer's token sell signature. Successfull call on this method grants, that the token sell can be registered by Replication Server",
	Flags:       verifyTokenSellSignatureFlags,
	Action:      actionDecorator(verifyTokenSellSignature),
}

const (
	flagIssuerTokenSellSignature = "issuer-signature"
)

var verifyTokenSellSignatureFlags = append(
	[]cli.Flag{
		cli.StringFlag{
			Name:  flagIssuerTokenSellSignature,
			Usage: "signature maden by responsible issuer, presenting issuer's good will to follow Sell terms",
		},
	},
	signTokenSaleFlags...,
)

func verifyTokenSellSignature(ctx *cli.Context) er.R {
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

	req := &replicator.VerifyTokenSellRequest{
		Sell: &replicator.TokenPurchase{
			Offer:           offer,
			IssuerSignature: issuerSignature,
		},
	}

	_, err = client.VerifyTokenSell(context.TODO(), req)
	if err != nil {
		return er.Errorf("requesting token sell signature verification: %s", err)
	}

	println("OK")

	return nil
}
