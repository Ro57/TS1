package signer

import (
	"fmt"

	"github.com/pkt-cash/pktd/lnd/chainreg"
	"github.com/pkt-cash/pktd/lnd/keychain"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/DB"
	hh "github.com/pkt-cash/pktd/lnd/lnrpc/tokens/hashHelper"
)

type Signer interface {
	Sign() (hh.Hash, error)
	Validate() (bool, error)
}

func NewSigner(v interface{}, s interface{}, cc *chainreg.ChainControl, nodeKeyDesc *keychain.KeyDescriptor) error {
	switch v := v.(type) {
	case DB.Block:
		signBlock := s.(*SignBlock)
		sign, err := NewSingBlock(v, cc, nodeKeyDesc)
		*signBlock = *sign

		return err
	default:
		return fmt.Errorf("type %T cannot be signed", v)

	}
}
