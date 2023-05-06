package keeper_test

import (
	abci "github.com/gridfx/fxchain/libs/tendermint/abci/types"

	"github.com/gridfx/fxchain/libs/cosmos-sdk/simapp"
	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/x/mint/internal/types"
)

// returns context and an app with updated mint keeper
func createTestApp(isCheckTx bool) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(isCheckTx)

	ctx := app.BaseApp.NewContext(isCheckTx, abci.Header{})
	app.MintKeeper.SetParams(ctx, types.DefaultParams())
	app.MintKeeper.SetMinter(ctx, types.InitialMinterCustom())

	return app, ctx
}
