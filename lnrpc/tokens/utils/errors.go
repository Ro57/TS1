package utils

import "github.com/pkt-cash/pktd/btcutil/er"

var (
	EmptyAddressErr = er.New("empty address")

	TokensDBNotFound       = er.New("tokens DB not created")
	InfoNotFoundErr        = er.New("token info not found")
	StateNotFoundErr       = er.New("state not found")
	TokenNotFoundErr       = er.New("token does not found")
	BlockNotFoundErr       = er.New("block info not found")
	RootHashNotFoundErr    = er.New("root hash not found")
	LastBlockNotFoundErr   = er.New("last block not found")
	ChainBucketNotFoundErr = er.New("chain bucket not found")
)
