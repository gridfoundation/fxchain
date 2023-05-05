package continuousauction

import (
	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"

	"github.com/gridfx/fxchain/x/order/keeper"
)

// nolint
type CaEngine struct {
}

// nolint
func (e *CaEngine) Run(ctx sdk.Context, keeper keeper.Keeper) {
	// TODO
}
