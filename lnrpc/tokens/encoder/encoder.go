package encoder

import "crypto/sha256"

type TokenSell struct {
	Token          string
	Price          uint64
	ValidUntilTime int64
	Count          uint64
}

func CreateHash(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)

	return hasher.Sum(nil)
}
