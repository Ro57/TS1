package replicatorrpc

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/DB"
	"github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/utils"
	"github.com/pkt-cash/pktd/pktwallet/walletdb"
)

func (s *Server) getToken(name string) (*replicator.Token, er.R) {
	var (
		token = &replicator.Token{
			Name:  name,
			Token: &DB.Token{},
		}
	)

	err := s.db.View(func(tx walletdb.ReadTx) er.R {
		rootBucket := tx.ReadBucket(utils.TokensKey)
		if rootBucket == nil {
			return utils.TokensDBNotFound
		}

		tokenBucket := rootBucket.NestedReadBucket([]byte(name))
		if tokenBucket == nil {
			return utils.TokenNotFoundErr
		}
		infoBytes := tokenBucket.Get(utils.InfoKey)
		if infoBytes == nil {
			return utils.InfoNotFoundErr
		}

		err := proto.Unmarshal(infoBytes, token.Token)
		if err != nil {
			return er.E(err)
		}

		token.Root = string(tokenBucket.Get(utils.RootHashKey))
		return nil
	})

	return token, err
}

func (s *Server) getIssuerTokens() (tokens IssuerTokens, err er.R) {
	err = s.db.View(func(tx walletdb.ReadTx) er.R {
		rootBucket := tx.ReadBucket(utils.TokensKey)
		if rootBucket == nil {
			return utils.TokensDBNotFound
		}

		issuerTokensBytes := rootBucket.Get(utils.IssuerTokens)
		if issuerTokensBytes == nil {
			tokens = IssuerTokens{}
			return nil
		}

		err := json.Unmarshal(issuerTokensBytes, &tokens)
		if err != nil {
			return er.E(err)
		}

		return nil
	})
	return
}
