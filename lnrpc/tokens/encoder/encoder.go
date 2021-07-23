package encoder

import "crypto/sha256"

type TokenSell struct {
	Token            string
	Price            uint64
	ID               string
	Identity_pubkey  string
	Host             string
	TokenHolderLogin string
	TokenBuyerLogin  string
	ValidUntilTime   int64
}

func CreateHash(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)

	return hasher.Sum(nil)
}
