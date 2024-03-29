package types

import (
	"github.com/gridfx/fxchain/libs/cosmos-sdk/codec"
	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"
	authtypes "github.com/gridfx/fxchain/libs/cosmos-sdk/x/auth/types"
	stakingtypes "github.com/gridfx/fxchain/libs/cosmos-sdk/x/staking/types"
)

// ModuleCdc defines a generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

// TODO: abstract genesis transactions registration back to staking
// required for genesis transactions
func init() {
	ModuleCdc = codec.New()
	stakingtypes.RegisterCodec(ModuleCdc)
	authtypes.RegisterCodec(ModuleCdc)
	sdk.RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
