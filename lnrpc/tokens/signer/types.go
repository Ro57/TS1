// extending the protobuf structure with signer methods
package signer

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/DB"
	hh "github.com/pkt-cash/pktd/lnd/lnrpc/tokens/hashHelper"
)

type SignBlock struct {
	// block — extended proobuf structure
	block DB.Block
	// block prefix used in start of hash string
	prefix string
}

func NewSingBlock(block DB.Block) (*SignBlock, error) {
	signBlock := &SignBlock{
		block:  block,
		prefix: "block",
	}

	_, err := signBlock.Sign()
	if err != nil {
		return nil, err
	}

	return signBlock, nil
}

var _ Signer = (*SignBlock)(nil)

// In this implementation generate sha256 hash
// TODO: get key for signing data and change algoritm
func (s *SignBlock) Sign() (hh.Hash, error) {
	s.block.Signature = ""

	buf, err := proto.Marshal(&s.block)

	if err != nil {
		return nil, fmt.Errorf("on marshal: %v", err)
	}

	sig := sha256.Sum256(buf)
	s.block.Signature = s.prefix + hex.EncodeToString(sig[:])

	blockSig, err := hh.NewBlock(s.block.Signature)
	if err != nil {
		return nil, err
	}
	return blockSig, nil
}

func (s *SignBlock) Validate() (bool, error) {
	sig := s.block.Signature
	s.block.Signature = ""

	buf, err := json.Marshal(s.block)
	if err != nil {
		return false, err
	}

	hash := sha256.Sum256(buf)
	blockHash := s.prefix + hex.EncodeToString(hash[:])
	s.block.Signature = sig

	return blockHash == sig, nil
}
