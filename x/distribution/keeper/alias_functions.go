package keeper

import (
	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/x/supply/exported"

	"github.com/gridfx/fxchain/x/distribution/types"
)

// GetDistributionAccount returns the distribution ModuleAccount
func (k Keeper) GetDistributionAccount(ctx sdk.Context) exported.ModuleAccountI {
	return k.supplyKeeper.GetModuleAccount(ctx, types.ModuleName)
}
