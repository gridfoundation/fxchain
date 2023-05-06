package types

import (
	abci "github.com/gridfx/fxchain/libs/tendermint/abci/types"
)

type CM40EventDataNewBlock struct {
	Block *CM40Block `json:"block"`

	ResultBeginBlock abci.ResponseBeginBlock `json:"result_begin_block"`
	ResultEndBlock   abci.ResponseEndBlock   `json:"result_end_block"`
}
