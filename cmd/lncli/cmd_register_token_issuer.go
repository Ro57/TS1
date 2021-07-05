package main

import (
	"context"
	"fmt"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"github.com/urfave/cli"
)

var registerTokenIssuerCommand = cli.Command{
	Name:     "registertokenissuer",
	Category: "Tokens",
	Usage:    "New token issuer registration",
	Description: `New token issuer registration
		
	You need to pass the username and password to register the new token issuer.
Token issuer login will be used to identify the session and execute command as authorized token holder`,

	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "login",
			Usage: "(required) Unique login by which storage session",
		},
		cli.StringFlag{
			Name:  "password",
			Usage: "(required) User password used for sign in",
		},
	},
	Action: registerTokenIssuer,
}

func registerTokenIssuer(ctx *cli.Context) er.R {
	client, cleanUp := getClient(ctx)
	defer cleanUp()

	login, err := parseRequiredString(ctx, "login")
	if err != nil {
		return er.E(err)
	}

	password, err := parseRequiredString(ctx, "password")
	if err != nil {
		return er.E(err)
	}

	// Request offers
	req := &replicator.RegisterRequest{
		Login:    login,
		Password: password,
	}

	_, err = client.RegisterTokenIssuer(context.TODO(), req)
	if err != nil {
		return er.E(err)
	}

	fmt.Println("Registration successful!")

	return nil
}
