package order

import (
	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"

	"github.com/gridfx/fxchain/x/common/perf"
	"github.com/gridfx/fxchain/x/order/keeper"
	"github.com/gridfx/fxchain/x/order/types"
	//"github.com/gridfx/fxchain/x/common/version"
)

// BeginBlocker runs the logic of BeginBlocker with version 0.
// BeginBlocker resets keeper cache.
func BeginBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	seq := perf.GetPerf().OnBeginBlockEnter(ctx, types.ModuleName)
	defer perf.GetPerf().OnBeginBlockExit(ctx, types.ModuleName, seq)

	keeper.ResetCache(ctx)
}
