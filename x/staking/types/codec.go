package types

import (
	"github.com/gridfx/fxchain/libs/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types for codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateValidator{}, "okfxchain/staking/MsgCreateValidator", nil)
	cdc.RegisterConcrete(MsgEditValidator{}, "okfxchain/staking/MsgEditValidator", nil)
	cdc.RegisterConcrete(MsgEditValidatorCommissionRate{}, "okfxchain/staking/MsgEditValidatorCommissionRate", nil)
	cdc.RegisterConcrete(MsgDestroyValidator{}, "okfxchain/staking/MsgDestroyValidator", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "okfxchain/staking/MsgDeposit", nil)
	cdc.RegisterConcrete(MsgWithdraw{}, "okfxchain/staking/MsgWithdraw", nil)
	cdc.RegisterConcrete(MsgAddShares{}, "okfxchain/staking/MsgAddShares", nil)
	cdc.RegisterConcrete(MsgRegProxy{}, "okfxchain/staking/MsgRegProxy", nil)
	cdc.RegisterConcrete(MsgBindProxy{}, "okfxchain/staking/MsgBindProxy", nil)
	cdc.RegisterConcrete(MsgUnbindProxy{}, "okfxchain/staking/MsgUnbindProxy", nil)
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
