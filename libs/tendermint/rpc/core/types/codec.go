package coretypes

import (
	amino "github.com/tendermint/go-amino"

	"github.com/gridfx/fxchain/libs/tendermint/types"
)

func RegisterAmino(cdc *amino.Codec) {
	types.RegisterEventDatas(cdc)
	types.RegisterBlockAmino(cdc)
}
