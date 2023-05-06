package connection

import (
	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"
	"github.com/gridfx/fxchain/libs/ibc-go/modules/core/03-connection/keeper"
	"github.com/gridfx/fxchain/libs/ibc-go/modules/core/03-connection/types"
)

// InitGenesis initializes the ibc connection submodule's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, gs types.GenesisState) {
	k.SetParams(ctx, gs.Params)
	for _, connection := range gs.Connections {
		conn := types.NewConnectionEnd(connection.State, connection.ClientId, connection.Counterparty, connection.Versions, connection.DelayPeriod)
		k.SetConnection(ctx, connection.Id, conn)
	}
	for _, connPaths := range gs.ClientConnectionPaths {
		k.SetClientConnectionPaths(ctx, connPaths.ClientId, connPaths.Paths)
	}
	k.SetNextConnectionSequence(ctx, gs.NextConnectionSequence)
	k.SetParams(ctx, gs.Params)
}

// ExportGenesis returns the ibc connection submodule's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.GenesisState{
		Connections:            k.GetAllConnections(ctx),
		ClientConnectionPaths:  k.GetAllClientConnectionPaths(ctx),
		NextConnectionSequence: k.GetNextConnectionSequence(ctx),
		Params:                 k.GetParams(ctx),
	}
}
