package main

import (
	"context"
	"time"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/issuer"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"github.com/urfave/cli"
)

// COMMENT ABOUT LOCKING FORM CALEB
// Perhaps we should use the LOCK_TOKEN justification hash as the ID to simplify validation. <— yes
// In this example, if after blocking 50 out of 100 tokens, someone wants to block more than 50 tokens, we can immediately refuse to block, since we will have data that the funds are already blocked for another user. <— yes, immediate rejection, no discredit needed
// This is a "simple" rejection
// Generally speaking, we should use "simple" rejection anywhere that the block can be known to be invalid only by looking at the chain. We should only use discredit in case of attacks such as duplicate block or fake timestamp
// This will very much reduce the risk of a discredit over a software bug or accident

var signTokenSaleCommand = cli.Command{
	Name:        "lock-tokens-for-transfer",
	Category:    "Tokens",
	Usage:       "Create token sell signature intent ",
	Description: "Create token sell signature intent by requesting open channel to sell from one user to another",
	Flags:       signTokenSaleFlags,
	Action:      actionDecorator(signTokenSell),
}

const (
	flagTokenBuyerLogin              = "token-buyer-login"
	flagTokenCount                   = "count"
	flagTokenName                    = "token"
	flagTokenPrice                   = "price"
	flagTokenHolderLogin             = "token-holder-login"
	flagIssuerOfferValidUntilSeconds = "issuer-offer-valid-until-seconds"
	flagIssuerID                     = "issuer-id"
	flagIssuerIdentityPubKey         = "issuer-identity-pubkey"
	flagIssuerHost                   = "issuer-host"
)

var signTokenSaleFlags = []cli.Flag{
	cli.StringFlag{
		Name:  flagTokenName,
		Usage: "target token identity to buy",
	},
	cli.Uint64Flag{
		Name:  flagTokenPrice,
		Usage: "token price per unit",
	},
	cli.Int64Flag{
		Name:  flagIssuerOfferValidUntilSeconds,
		Usage: "issuer's token offer term's time constraint",
	},
	cli.StringFlag{
		Name:  flagIssuerIdentityPubKey,
		Usage: "issuer node's pubkey to initialize a channel",
	},
	cli.Int64Flag{
		Name:  flagTokenCount,
		Usage: "number of tokens issued",
	},
}

func signTokenSell(ctx *cli.Context) er.R {
	client, cleanUp := getClient(ctx)
	defer cleanUp()

	offer, _err := extractTokenSell(ctx)
	if _err != nil {
		return _err
	}

	signTokenSellReq := &issuer.SignTokenSellRequest{
		Offer: offer,
	}

	signTokenSellResponse, err := client.SignTokenSell(context.TODO(), signTokenSellReq)
	if err != nil {
		return er.Errorf("requesting token purchase signature: %s", err)
	}

	printRespJSON(signTokenSellResponse)

	return nil
}

func extractTokenSell(ctx *cli.Context) (*replicator.TokenOffer, er.R) {
	token, err := parseRequiredString(ctx, flagTokenName)
	if err != nil {
		return nil, er.E(err)
	}

	price, err := parseRequiredUint64(ctx, flagTokenPrice)
	if err != nil {
		return nil, er.E(err)
	}

	validUntilSeconds, err := parseRequiredInt64(ctx, flagIssuerOfferValidUntilSeconds)
	if err != nil {
		return nil, er.E(err)
	}

	count, err := parseRequiredUint64(ctx, flagTokenCount)
	if err != nil {
		return nil, er.E(err)
	}

	if expDate := time.Unix(validUntilSeconds, 0).UTC(); expDate.Before(time.Now().UTC()) {
		return nil, er.Errorf("%q argument provided is in the past or empty", flagIssuerOfferValidUntilSeconds)
	}

	// Extract general token offer data
	offer := &replicator.TokenOffer{
		Token:             token,
		Price:             price,
		ValidUntilSeconds: validUntilSeconds,
		Count:             count,
	}

	return offer, nil
}
