package main

import (
	"context"
	"fmt"
	"time"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"github.com/urfave/cli"
)

var issueTokenCommand = cli.Command{
	Name:        "issuetoken",
	Category:    "Tokens",
	Usage:       "Issue new token",
	Description: "Issue new token. This command is only allowed users with issuer role",
	Flags:       issueTokenFlags,
	Action:      actionDecorator(issueToken),
}

const (
	flagTokenIssuerLogin = "token-issuer-login"
)

var issueTokenFlags = []cli.Flag{
	cli.StringFlag{
		Name:  flagTokenName,
		Usage: "target token identity to buy",
	},
	cli.Uint64Flag{
		Name:  flagTokenPrice,
		Usage: "token price per unit",
	},
	cli.StringFlag{
		Name:  flagTokenIssuerLogin,
		Usage: "issuer identity, a token createor",
	},
	cli.Int64Flag{
		Name:  flagIssuerOfferValidUntilSeconds,
		Usage: "issuer's token offer term's time constraint",
	},
	cli.StringFlag{
		Name:  flagIssuerID,
		Usage: "issuer identity to buy target token from",
	},
}

func issueToken(ctx *cli.Context) er.R {
	client, cleanUp := getClient(ctx)
	defer cleanUp()

	offer, _err := extractTokenIssue(ctx)
	if _err != nil {
		return _err
	}

	issuerTokenReq := &replicator.IssueTokenRequest{
		Offer: offer,
	}
	_, err := client.IssueToken(context.TODO(), issuerTokenReq)
	if err != nil {
		return er.Errorf("requesting token issue: %s", err)
	}

	fmt.Println("issue succesful!")

	return nil
}

func extractTokenIssue(ctx *cli.Context) (*replicator.TokenOffer, er.R) {
	// Extract general token offer data
	offer := &replicator.TokenOffer{}

	offer.Token = ctx.String(flagTokenName)
	if offer.Token == "" {
		return nil, er.Errorf("empty %q argument provided", flagTokenName)
	}
	offer.Price = ctx.Uint64(flagTokenPrice)
	if offer.Price == 0 {
		return nil, er.Errorf("empty %q argument provided", flagTokenPrice)
	}
	offer.TokenHolderLogin = ctx.String(flagTokenIssuerLogin)
	if offer.TokenHolderLogin == "" {
		return nil, er.Errorf("empty %q argument provided", flagTokenIssuerLogin)
	}

	offer.ValidUntilSeconds = ctx.Int64(flagIssuerOfferValidUntilSeconds)

	if expDate := time.Unix(offer.ValidUntilSeconds, 0).UTC(); expDate.Before(time.Now().UTC()) {
		return nil, er.Errorf("%q argument provided is in the past or empty", flagIssuerOfferValidUntilSeconds)
	}

	// Extract token offer issuer data
	issuerInfo := &replicator.IssuerInfo{}

	issuerInfo.Id = ctx.String(flagIssuerID)
	if issuerInfo.Id == "" {
		return nil, er.Errorf("empty %q argument provided", flagIssuerID)
	}

	offer.IssuerInfo = issuerInfo

	return offer, nil
}
