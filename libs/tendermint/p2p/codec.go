package p2p

import (
	amino "github.com/tendermint/go-amino"

	cryptoamino "github.com/gridfx/fxchain/libs/tendermint/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoamino.RegisterAmino(cdc)
}
