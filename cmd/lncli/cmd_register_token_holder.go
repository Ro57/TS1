package main

import (
	"context"
	"fmt"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/replicator"
	"github.com/urfave/cli"
)

var registerTokenHolderCoomand = cli.Command{
	Name:     "registertokenholder",
	Category: "Tokens",
	Usage:    "New token holder registration",
	Description: `New token holder registration
		
	You need to pass the username and password to register the new token holder.
Token holder login will be used to identify the session and execute command as authorized token holder`,

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
	Action: registerTokenHolder,
}

func registerTokenHolder(ctx *cli.Context) er.R {
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

	_, err = client.RegisterTokenHolder(context.TODO(), req)
	if err != nil {
		return er.E(err)
	}

	fmt.Println("Registration successful!")

	return nil
}
