package main

import (
	"context"
	"time"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/issuer"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/replicator"
	"github.com/urfave/cli"
)

var signTokenPurchaseCommand = cli.Command{
	Name:        "signtokenpurchase",
	Category:    "Tokens",
	Usage:       "Signs token purchase intent",
	Description: "Signs token purchase intent by requesting Issuer's signature",
	Flags:       signTokenPurchaseFlags,
	Action:      actionDecorator(signTokenPurchase),
}

const (
	flagTokenName                    = "token"
	flagTokenPrice                   = "price"
	flagTokenHolderLogin             = "token-holder-login"
	flagIssuerOfferValidUntilSeconds = "issuer-offer-valid-until-seconds"
	flagIssuerID                     = "issuer-id"
	flagIssuerIdentityPubKey         = "issuer-identity-pubkey"
	flagIssuerHost                   = "issuer-host"
)

var signTokenPurchaseFlags = []cli.Flag{
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

func signTokenPurchase(ctx *cli.Context) er.R {
	client, cleanUp := getClient(ctx)
	defer cleanUp()

	offer, _err := extractTokenOffer(ctx)
	if _err != nil {
		return _err
	}

	signTokenPurchaseReq := &issuer.SignTokenPurchaseRequest{
		Offer: offer,
	}
	signTokenPurchaseResp, err := client.SignTokenPurchase(context.TODO(), signTokenPurchaseReq)
	if err != nil {
		return er.Errorf("requesting token purchase signature: %s", err)
	}

	printRespJSON(signTokenPurchaseResp)

	return nil
}

func extractTokenOffer(ctx *cli.Context) (*replicator.TokenOffer, er.R) {
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
	offer.TokenHolderLogin = ctx.String(flagTokenHolderLogin)
	if offer.TokenHolderLogin == "" {
		return nil, er.Errorf("empty %q argument provided", flagTokenHolderLogin)
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
	issuerInfo.IdentityPubkey = ctx.String(flagIssuerIdentityPubKey)
	if issuerInfo.IdentityPubkey == "" {
		return nil, er.Errorf("empty %q argument provided", flagIssuerIdentityPubKey)
	}
	issuerInfo.Host = ctx.String(flagIssuerHost)
	if issuerInfo.Host == "" {
		return nil, er.Errorf("empty %q argument provided", flagIssuerHost)
	}

	offer.IssuerInfo = issuerInfo

	return offer, nil
}
