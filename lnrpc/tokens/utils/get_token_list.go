package utils

import (
	"github.com/golang/protobuf/proto"
	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/DB"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/tokendb"
	"github.com/pkt-cash/pktd/pktwallet/walletdb"
)

var (
	IssuerTokens = []byte("issuer_tokens")

	InfoKey   = []byte("info")
	StateKey  = []byte("state")
	ChainKey  = []byte("chain")
	TokensKey = []byte("tokens")
	// rootHash is a hash of last block in chain
	RootHashKey = []byte("rootHash")
	// Replication is a information about replication server configuration
	Replication = []byte("replication")
	// Issuer is a collection with key issuer pubKey
	Issuers = []byte("issuers")
)

func GetTokenList(db *tokendb.TokenStrikeDB) ([]*replicator.Token, error) {
	var resultList []*replicator.Token

	err := db.View(func(tx walletdb.ReadTx) er.R {
		rootBucket := tx.ReadBucket(TokensKey)
		if rootBucket == nil {
			return er.New("tokens do not exist")
		}

		return rootBucket.ForEach(func(k, _ []byte) er.R {
			tokenBucket := rootBucket.NestedReadBucket(k)

			var dbToken DB.Token
			err := proto.Unmarshal(tokenBucket.Get(InfoKey), &dbToken)
			if err != nil {
				return er.E(err)
			}

			token := replicator.Token{
				Name:  string(k),
				Token: &dbToken,
				Root:  string(tokenBucket.Get(RootHashKey)),
			}

			resultList = append(resultList, &token)
			return nil
		})
	})
	if err != nil {
		return nil, err.Native()
	}

	return resultList, nil
}
