package main

import (
	"context"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/issuer"
	"github.com/urfave/cli"
)

// COMMENT ABOUT LOCKING FORM CALEB
// Perhaps we should use the LOCK_TOKEN justification hash as the ID to simplify validation. <— yes
// In this example, if after blocking 50 out of 100 tokens, someone wants to block more than 50 tokens, we can immediately refuse to block, since we will have data that the funds are already blocked for another user. <— yes, immediate rejection, no discredit needed
// This is a "simple" rejection
// Generally speaking, we should use "simple" rejection anywhere that the block can be known to be invalid only by looking at the chain. We should only use discredit in case of attacks such as duplicate block or fake timestamp
// This will very much reduce the risk of a discredit over a software bug or accident

var lockTokenCommand = cli.Command{
	Name:     "lock-tokens-for-transfer",
	Category: "Tokens",
	Usage:    "Create new lock",
	Description: `Create new lock with information about htlc preimagine and 
	return lokc ID as a hash of lock structure`,
	Flags:  lockTokenFlags,
	Action: actionDecorator(lockToken),
}

const (
	flagTokenName     = "token"
	flagHTLC          = "htlc"
	flagRecipient     = "address"
	defaultProofCount = 100
)

var lockTokenFlags = []cli.Flag{
	cli.StringFlag{
		Name:  flagTokenName,
		Usage: "target token to buy",
	},
	cli.Uint64Flag{
		Name:  flagCount,
		Usage: "count of locked token",
	},
	cli.StringFlag{
		Name:  flagHTLC,
		Usage: "hash of htcl",
	},
	cli.StringFlag{
		Name:  flagRecipient,
		Usage: "address of wallet to receive payment",
	},
}

func lockToken(ctx *cli.Context) er.R {
	client, cleanUp := getClient(ctx)
	defer cleanUp()

	lockTokenRequest, _err := extractTokenLock(ctx)
	if _err != nil {
		return _err
	}

	lockTokenResponse, err := client.LockToken(context.TODO(), lockTokenRequest)
	if err != nil {
		return er.Errorf("requesting token purchase signature: %s", err)
	}

	printRespJSON(lockTokenResponse)

	return nil
}

func extractTokenLock(ctx *cli.Context) (*issuer.LockTokenRequest, er.R) {
	token, err := parseRequiredString(ctx, flagTokenName)
	if err != nil {
		return nil, er.E(err)
	}

	count, err := parseRequiredInt64(ctx, flagCount)
	if err != nil {
		return nil, er.E(err)
	}

	htlc, err := parseRequiredString(ctx, flagHTLC)
	if err != nil {
		return nil, er.E(err)
	}

	recipient, err := parseRequiredString(ctx, flagRecipient)
	if err != nil {
		return nil, er.E(err)
	}

	// Extract general token offer data
	offer := &issuer.LockTokenRequest{
		Token:      token,
		Count:      count,
		Htlc:       htlc,
		Recipient:  recipient,
		ProofCount: defaultProofCount,
	}

	return offer, nil
}
