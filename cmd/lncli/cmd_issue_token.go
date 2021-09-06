package main

import (
	"context"
	"fmt"
	"time"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/DB"
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
	cli.Int64Flag{
		Name:  flagCount,
		Usage: "number of issued token",
	},
	cli.Int64Flag{
		Name:  flagExpirationBlockNumber,
		Usage: "number of PKT block after which the token expires",
	},
	cli.StringFlag{
		Name:  flagUrl,
		Usage: "urls for access to blockchain",
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

	fmt.Println("issue successful!")

	return nil
}

func extractTokenIssue(ctx *cli.Context) (*DB.Token, er.R) {
	// Extract general token offer data
	offer := &DB.Token{}

	offer.Count = ctx.Int64(flagCount)
	if offer.Count == 0 {
		return nil, er.Errorf("empty %q argument provided", flagCount)
	}

	offer.Expiration = int32(ctx.Int64(flagExpirationBlockNumber))
	if offer.Expiration == 0 {
		return nil, er.Errorf("empty %q argument provided", flagExpirationBlockNumber)
	}

	offer.Url = ctx.String(flagUrl)
	if offer.Url == "" {
		return nil, er.Errorf("empty %q argument provided", flagUrl)
	}

	offer.Creation = time.Now().Unix()

	return offer, nil
}
