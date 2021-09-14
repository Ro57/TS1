package main

import (
	"context"
	"fmt"
	"time"

	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/DB"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"github.com/urfave/cli"
)

var issueTokenCommand = cli.Command{
	Name:        "issue-token",
	Category:    "Tokens",
	Usage:       "Issue new token",
	Description: "Issue new token. This command is only allowed users with issuer role",
	Flags:       issueTokenFlags,
	Action:      actionDecorator(issueToken),
}

const (
	flagUrl                   = "url"
	flagCount                 = "count"
	flagExpirationBlockNumber = "expiration-block-number"
)

var issueTokenFlags = []cli.Flag{
	cli.StringFlag{
		Name:  flagTokenName,
		Usage: "target token identity to issue",
	},
	cli.Int64Flag{
		Name:  flagCount,
		Usage: "number of issued token",
	},
	cli.Int64Flag{
		Name:  flagExpirationBlockNumber,
		Usage: "number of PKT block after which the token expires",
	},
	cli.StringSliceFlag{
		Name:  flagUrl,
		Usage: "urls for access to blockchain",
	},
}

func issueToken(ctx *cli.Context) er.R {
	client, cleanUp := getClient(ctx)
	defer cleanUp()

	issuerTokenReq, _err := extractTokenIssue(ctx)
	if _err != nil {
		return _err
	}

	_, err := client.IssueToken(context.TODO(), issuerTokenReq)
	if err != nil {
		return er.Errorf("requesting token issue: %s", err)
	}

	fmt.Println("issue successful!")

	return nil
}

func extractTokenIssue(ctx *cli.Context) (*replicator.IssueTokenRequest, er.R) {
	// Extract general token offer data
	offer := &replicator.IssueTokenRequest{
		Name:  "",
		Offer: &DB.Token{},
	}

	offer.Name = ctx.String(flagTokenName)
	if offer.Name == "" {
		return nil, er.Errorf("empty %q argument provided", flagTokenName)
	}

	offer.Offer.Count = ctx.Int64(flagCount)
	if offer.Offer.Count == 0 {
		return nil, er.Errorf("empty %q argument provided", flagCount)
	}

	offer.Offer.Expiration = int32(ctx.Int64(flagExpirationBlockNumber))
	if offer.Offer.Expiration == 0 {
		return nil, er.Errorf("empty %q argument provided", flagExpirationBlockNumber)
	}

	offer.Offer.Urls = ctx.StringSlice(flagUrl)
	if len(offer.Offer.Urls) == 0 {
		return nil, er.Errorf("empty %q argument provided", flagUrl)
	}

	offer.Offer.Creation = time.Now().Unix()

	return offer, nil
}
