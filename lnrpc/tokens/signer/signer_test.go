package signer

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/pkt-cash/pktd/lnd/chainreg"
	"github.com/pkt-cash/pktd/lnd/keychain"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/DB"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/justifications"
)

func TestSigner(t *testing.T) {
	justificationPool := []*DB.Justification{}

	justificationPool = append(justificationPool,
		&DB.Justification{
			Content: &DB.Justification_Transfer{
				Transfer: &justifications.TranferToken{
					HtlcSecret: "some",
					Lock:       "some",
				},
			},
		},
	)

	wantBlock := DB.Block{
		Justifications: justificationPool,
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

	activeChainControl := &chainreg.ChainControl{}
	fmt.Print("cc created")
	idKeyDesc, errr := activeChainControl.KeyRing.DeriveKey(
		keychain.KeyLocator{
			Family: keychain.KeyFamilyNodeKey,
			Index:  0,
		},
	)
	if errr != nil {
		t.Fatalf("generate key descriptor: %v", errr.Native())
	}

	var sign SignBlock

	err := NewSigner(wantBlock, &sign, activeChainControl, &idKeyDesc)
	if err != nil {
		t.Fatalf("on create signer %v", err)
	}

	wantSignBlock, err := NewSingBlock(wantBlock, activeChainControl, &idKeyDesc)
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
		t.Fatalf("want prefix %s but get %s", wantSigner.prefix, sign.prefix)
	}
}
