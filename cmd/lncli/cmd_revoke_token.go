package main

import (
	"context"
	"fmt"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc"
	"github.com/urfave/cli"
)

var revokeTokenCommand = cli.Command{
	Name:        "revoketoken",
	Category:    "Tokens",
	Usage:       "Revoke token from token pull",
	Description: "Revoke token from token pull. Revoke may be execute only token issuer",
	Flags:       revokeTokenFlags,
	Action:      actionDecorator(revokeToken),
}

var revokeTokenFlags = []cli.Flag{
	cli.StringFlag{
		Name:  flagTokenName,
		Usage: "target token identity to revoke",
	},
	cli.StringFlag{
		Name:  "login",
		Usage: "issuer login, used to verify identity",
	},

	cli.StringFlag{
		Name:  flagIssuerHost,
		Usage: "issuer node's host to initialize a channel",
	},
}

func revokeToken(ctx *cli.Context) er.R {
	client, cleanUp := getClient(ctx)
	defer cleanUp()

	tokenName, err := parseRequiredString(ctx, flagTokenName)
	if err != nil {
		return er.E(err)
	}

	issuerHost, err := parseRequiredString(ctx, flagIssuerHost)
	if err != nil {
		return er.E(err)
	}

	login, err := parseRequiredString(ctx, "login")
	if err != nil {
		return er.E(err)
	}

	revokeTokenReq := &lnrpc.RevokeTokenRequest{
		TokenName:  tokenName,
		Login:      login,
		IssuerHost: issuerHost,
	}

	_, err = client.RevokeToken(context.TODO(), revokeTokenReq)
	if err != nil {
		return er.Errorf("requesting token revoke: %s", err)
	}

	fmt.Println("revoke succesful!")

	return nil
}
