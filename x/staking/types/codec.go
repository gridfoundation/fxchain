package types

import (
	"github.com/gridfx/fxchain/libs/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types for codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateValidator{}, "fxchain/staking/MsgCreateValidator", nil)
	cdc.RegisterConcrete(MsgEditValidator{}, "fxchain/staking/MsgEditValidator", nil)
	cdc.RegisterConcrete(MsgEditValidatorCommissionRate{}, "fxchain/staking/MsgEditValidatorCommissionRate", nil)
	cdc.RegisterConcrete(MsgDestroyValidator{}, "fxchain/staking/MsgDestroyValidator", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "fxchain/staking/MsgDeposit", nil)
	cdc.RegisterConcrete(MsgWithdraw{}, "fxchain/staking/MsgWithdraw", nil)
	cdc.RegisterConcrete(MsgAddShares{}, "fxchain/staking/MsgAddShares", nil)
	cdc.RegisterConcrete(MsgRegProxy{}, "fxchain/staking/MsgRegProxy", nil)
	cdc.RegisterConcrete(MsgBindProxy{}, "fxchain/staking/MsgBindProxy", nil)
	cdc.RegisterConcrete(MsgUnbindProxy{}, "fxchain/staking/MsgUnbindProxy", nil)
	cdc.RegisterConcrete(CM45Validator{}, "cosmos-sdk/staking/validator", nil)
}

// ModuleCdc is generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
