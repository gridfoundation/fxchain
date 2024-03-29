package keeper

import (
	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"
	authexported "github.com/gridfx/fxchain/libs/cosmos-sdk/x/auth/exported"
	evmtypes "github.com/gridfx/fxchain/x/evm/types"
	wasmtypes "github.com/gridfx/fxchain/x/wasm/types"
)

type EVMKeeper interface {
	GetChainConfig(ctx sdk.Context) (evmtypes.ChainConfig, bool)
	GenerateCSDBParams() evmtypes.CommitStateDBParams
	GetParams(ctx sdk.Context) evmtypes.Params
}

type WASMKeeper interface {
	// Execute executes the contract instance
	Execute(ctx sdk.Context, contractAddress sdk.WasmAddress, caller sdk.WasmAddress, msg []byte, coins sdk.Coins) ([]byte, error)
	GetParams(ctx sdk.Context) wasmtypes.Params
}

// AccountKeeper defines the expected account keeper interface
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authexported.Account
	SetAccount(ctx sdk.Context, acc authexported.Account)
	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) authexported.Account
}

type BankKeeper interface {
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
}
