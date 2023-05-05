package auth

import (
	"github.com/gridfx/fxchain/libs/cosmos-sdk/x/auth/exported"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/x/auth/keeper"
)

type (
	Account       = exported.Account
	ModuleAccount = exported.ModuleAccount
	ObserverI     = keeper.ObserverI
)
