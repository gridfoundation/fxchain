package client

import (
	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"
	sdkerrors "github.com/gridfx/fxchain/libs/cosmos-sdk/types/errors"
	"github.com/gridfx/fxchain/libs/ibc-go/modules/core/02-client/keeper"
	"github.com/gridfx/fxchain/libs/ibc-go/modules/core/02-client/types"
	govtypes "github.com/gridfx/fxchain/x/gov/types"
)

// NewClientUpdateProposalHandler defines the client update proposal handler
func NewClientUpdateProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content *govtypes.Proposal) sdk.Error {
		cont := content.Content
		switch c := cont.(type) {
		case *types.ClientUpdateProposal:
			return k.ClientUpdateProposal(ctx, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized ibc proposal content type: %T", c)
		}
	}
}
