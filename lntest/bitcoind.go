// +build bitcoind
// +build !notxindex

package lntest

import (
	"github.com/pkt-cash/pktd/chaincfg"
)

// NewBackend starts a bitcoind node with the txindex enabled and returns a
// BitcoindBackendConfig for that node.
func NewBackend(miner string, netParams *chaincfg.Params) (
	*BitcoindBackendConfig, func() error, er.R) {

	extraArgs := []string{
		"-debug",
		"-regtest",
		"-txindex",
		"-disablewallet",
	}

	return newBackend(miner, netParams, extraArgs)
}
