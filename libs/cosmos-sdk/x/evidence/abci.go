package evidence

import (
	"fmt"

	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"

	abci "github.com/gridfx/fxchain/libs/tendermint/abci/types"
	tmtypes "github.com/gridfx/fxchain/libs/tendermint/types"
)

// BeginBlocker iterates through and handles any newly discovered evidence of
// misbehavior submitted by Tendermint. Currently, only equivocation is handled.
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k Keeper) {
	for _, tmEvidence := range req.ByzantineValidators {
		switch tmEvidence.Type {
		case tmtypes.ABCIEvidenceTypeDuplicateVote:
			evidence := ConvertDuplicateVoteEvidence(tmEvidence)
			k.HandleDoubleSign(ctx, evidence.(Equivocation))

		default:
			k.Logger(ctx).Error(fmt.Sprintf("ignored unknown evidence type: %s", tmEvidence.Type))
		}
	}
}
