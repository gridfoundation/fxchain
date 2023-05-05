package keeper

import (
	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"

	"github.com/gridfx/fxchain/x/distribution/types"
	"github.com/gridfx/fxchain/x/staking/exported"
)

// initialize rewards for a new validator
func (k Keeper) initializeValidator(ctx sdk.Context, val exported.ValidatorI) {
	if k.CheckDistributionProposalValid(ctx) {
		k.initializeValidatorDistrProposal(ctx, val)
		return
	}

	// set accumulated commissions
	k.SetValidatorAccumulatedCommission(ctx, val.GetOperator(), types.InitialValidatorAccumulatedCommission())
}
