package transfer

import (
	"github.com/gridfx/fxchain/libs/ibc-go/modules/apps/transfer/keeper"
	"github.com/gridfx/fxchain/libs/ibc-go/modules/apps/transfer/types"
)

var (
	NewKeeper  = keeper.NewKeeper
	ModuleCdc  = types.ModuleCdc
	SetMarshal = types.SetMarshal
	NewQuerier = keeper.NewQuerier
)
