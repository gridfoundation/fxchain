package keeper

import (
	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"
	"github.com/gridfx/fxchain/x/evm/types"
)

// GetParams returns the total set of evm parameters.
func (k *Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	if ctx.UseParamCache() {
		if types.GetEvmParamsCache().IsNeedParamsUpdate() {
			params = k.getParams(ctx)
			types.GetEvmParamsCache().UpdateParams(params, ctx.IsCheckTx())
		} else {
			params = types.GetEvmParamsCache().GetParams()
		}
	} else {
		params = k.getParams(ctx)
	}

	return
}

func (k *Keeper) getParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return
}

// SetParams sets the evm parameters to the param space.
func (k *Keeper) SetParams(ctx sdk.Context, params types.Params) {
	if k.EvmStateDb != nil {
		k.EvmStateDb.WithContext(ctx).SetParams(params)
	}

	k.paramSpace.SetParamSet(ctx, &params)
	types.GetEvmParamsCache().SetNeedParamsUpdate()
}
