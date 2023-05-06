package keeper

import (
	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"
	govtypes "github.com/gridfx/fxchain/x/gov/types"
)

// GovKeeper defines the expected gov Keeper
type GovKeeper interface {
	GetDepositParams(ctx sdk.Context) govtypes.DepositParams
	GetVotingParams(ctx sdk.Context) govtypes.VotingParams
}
