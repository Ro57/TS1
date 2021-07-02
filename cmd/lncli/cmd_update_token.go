package main

import (
	"context"
	"fmt"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/issuer"
	"github.com/urfave/cli"
)

var updateTokenCommand = cli.Command{
	Name:     "updatetoken",
	Category: "Tokens",
	Usage:    "Update inforamtion about token",
	Description: `Update inforamtion about token. May be execute 
only issuer of this token, and to the already existing token.`,
	Flags:  updateTokenFlags,
	Action: actionDecorator(updateToken),
}

var updateTokenFlags = issueTokenFlags

func updateToken(ctx *cli.Context) er.R {
	client, cleanUp := getClient(ctx)
	defer cleanUp()

	offer, _err := extractTokenIssue(ctx)
	if _err != nil {
		return _err
	}

	updateTokenReq := &issuer.UpdateTokenRequest{
		Offer: offer,
	}

	_, err := client.UpdateToken(context.TODO(), updateTokenReq)
	if err != nil {
		return er.Errorf("requesting token inforamtion update: %s", err)
	}

	fmt.Println("update succesful!")

	return nil
}
