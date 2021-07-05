package main

import (
	"context"
	"time"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/issuer"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"github.com/urfave/cli"
)

var signTokenSaleCommand = cli.Command{
	Name:        "maketokensellsignature",
	Category:    "Tokens",
	Usage:       "Create token sell signature intent ",
	Description: "Create token sell signature intent by requesting open channel to sell from one user to another",
	Flags:       signTokenSaleFlags,
	Action:      actionDecorator(signTokenSell),
}

const (
	flagTokenBuyerLogin = "token-buyer-login"
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
	cli.StringFlag{
		Name:  flagTokenHolderLogin,
		Usage: "seller identity, a current token holder",
	},
	cli.StringFlag{
		Name:  flagTokenBuyerLogin,
		Usage: "buyer identity, a future bought token holder",
	},
	cli.Int64Flag{
		Name:  flagIssuerOfferValidUntilSeconds,
		Usage: "issuer's token offer term's time constraint",
	},
	cli.StringFlag{
		Name:  flagIssuerID,
		Usage: "issuer identity to buy target token from",
	},
	cli.StringFlag{
		Name:  flagIssuerIdentityPubKey,
		Usage: "issuer node's pubkey to initialize a channel",
	},
	cli.StringFlag{
		Name:  flagIssuerHost,
		Usage: "issuer node's host to initialize a channel",
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

	signTokenPurchaseResp, err := client.SignTokenSell(context.TODO(), signTokenSellReq)
	if err != nil {
		return er.Errorf("requesting token purchase signature: %s", err)
	}

	printRespJSON(signTokenPurchaseResp)

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

	tokenHolderLogin, err := parseRequiredString(ctx, flagTokenHolderLogin)
	if err != nil {
		return nil, er.E(err)
	}

	tokenBuyerLogin, err := parseRequiredString(ctx, flagTokenBuyerLogin)
	if err != nil {
		return nil, er.E(err)
	}

	validUntilSeconds, err := parseRequiredInt64(ctx, flagIssuerOfferValidUntilSeconds)
	if err != nil {
		return nil, er.E(err)
	}

	if expDate := time.Unix(validUntilSeconds, 0).UTC(); expDate.Before(time.Now().UTC()) {
		return nil, er.Errorf("%q argument provided is in the past or empty", flagIssuerOfferValidUntilSeconds)
	}

	ID, err := parseRequiredString(ctx, flagIssuerID)
	if err != nil {
		return nil, er.E(err)
	}

	identityPubkey, err := parseRequiredString(ctx, flagIssuerIdentityPubKey)
	if err != nil {
		return nil, er.E(err)
	}

	host, err := parseRequiredString(ctx, flagIssuerHost)
	if err != nil {
		return nil, er.E(err)
	}

	// Extract general token offer data
	offer := &replicator.TokenOffer{
		Token:             token,
		Price:             price,
		TokenHolderLogin:  tokenHolderLogin,
		TokenBuyerLogin:   tokenBuyerLogin,
		ValidUntilSeconds: validUntilSeconds,

		// Extract token offer issuer data
		IssuerInfo: &replicator.IssuerInfo{
			Id:             ID,
			IdentityPubkey: identityPubkey,
			Host:           host,
		},
	}

	return offer, nil
}
