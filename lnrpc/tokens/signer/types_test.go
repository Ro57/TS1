package signer

import (
	"testing"
	"time"

	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/DB"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/justifications"
	hashhelper "github.com/pkt-cash/pktd/lnd/lnrpc/tokens/hashHelper"
)

func TestBlock(t *testing.T) {
	wantBlock := DB.Block{
		Justification: &DB.Block_Transfer{
			Transfer: &justifications.TranferToken{
				HtlcSecretHash: "some",
				Lock:           "some",
			},
		},
		PrevBlock:      "hashPrevBlock",
		Creation:       time.Now().Unix(),
		State:          "hashOfState",
		PktBlockHash:   "hashFromPkt",
		PktBlockHeight: 1000,
		Height:         10,
	}

	signBlock, err := NewSingBlock(wantBlock)
	if err != nil {
		t.Fatalf("Create singed block: %v", err)
	}
	if signBlock.block.Signature == "" {
		t.Fatal("Empty signature")
	}

	wantHash, err := hashhelper.NewBlock(signBlock.block.Signature)
	if err != nil {
		t.Fatalf("generate block from hashhelper: %v", err)
	}

	blockHash, err := signBlock.Sign()
	if err != nil {
		t.Fatalf("Sing error: %v", err)
	}
	if !blockHash.Validate() {
		t.Fatalf("Incorrect string format of %v", blockHash)
	}
	if *blockHash.(*hashhelper.Block) != *wantHash {
		t.Fatalf("get hash: %v \n want hash: %v", *blockHash.(*hashhelper.Block), *wantHash)
	}
}
