package icamauth

import (
	"fmt"

	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"
	sdkerrors "github.com/gridfx/fxchain/libs/cosmos-sdk/types/errors"
	tmtypes "github.com/gridfx/fxchain/libs/tendermint/types"
	"github.com/gridfx/fxchain/x/icamauth/types"
)

// NewHandler returns sdk.Handler for IBC token transfer module messages
func NewHandler(k types.MsgServer) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		if !tmtypes.HigherThanVenus4(ctx.BlockHeight()) {
			errMsg := fmt.Sprintf("icamauth  is not supported at height %d", ctx.BlockHeight())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}

		ctx.SetEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgSubmitTx:
			res, err := k.SubmitTx(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRegisterAccount:
			res, err := k.RegisterAccount(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized ICS-20 transfer message type: %T", msg)
		}
	}
}
