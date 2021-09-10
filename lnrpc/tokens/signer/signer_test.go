package signer

import (
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/DB"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/justifications"
)

func TestSigner(t *testing.T) {
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

	wantSigner := SignBlock{
		block:  wantBlock,
		prefix: "block",
	}

	var sign SignBlock

	err := NewSigner(wantBlock, &sign)
	if err != nil {
		t.Fatalf("on create signer %v", err)
	}

	wantSignBlock, err := NewSingBlock(wantBlock)
	if err != nil {
		t.Fatalf("create sign block from want block: %v", err)
	}

	bytesWantBlock, err := proto.Marshal(&wantSignBlock.block)
	if err != nil {
		t.Fatalf("marshal want block: %v", err)
	}

	bytesSignBlock, err := proto.Marshal(&sign.block)
	if err != nil {
		t.Fatalf("marshal sing block: %v", err)
	}

	if string(bytesWantBlock) != string(bytesSignBlock) {
		t.Fatalf("want \n %v block but get \n %v", bytesWantBlock, bytesSignBlock)
	}

	if sign.prefix != wantSigner.prefix {
		t.Fatalf("want prefix %s but get %v", wantSigner.prefix, sign)
	}
}
