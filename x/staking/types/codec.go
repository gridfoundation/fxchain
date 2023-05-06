package types

import (
	"github.com/gridfx/fxchain/libs/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types for codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateValidator{}, "gridfxchain/staking/MsgCreateValidator", nil)
	cdc.RegisterConcrete(MsgEditValidator{}, "gridfxchain/staking/MsgEditValidator", nil)
	cdc.RegisterConcrete(MsgEditValidatorCommissionRate{}, "gridfxchain/staking/MsgEditValidatorCommissionRate", nil)
	cdc.RegisterConcrete(MsgDestroyValidator{}, "gridfxchain/staking/MsgDestroyValidator", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "gridfxchain/staking/MsgDeposit", nil)
	cdc.RegisterConcrete(MsgWithdraw{}, "gridfxchain/staking/MsgWithdraw", nil)
	cdc.RegisterConcrete(MsgAddShares{}, "gridfxchain/staking/MsgAddShares", nil)
	cdc.RegisterConcrete(MsgRegProxy{}, "gridfxchain/staking/MsgRegProxy", nil)
	cdc.RegisterConcrete(MsgBindProxy{}, "gridfxchain/staking/MsgBindProxy", nil)
	cdc.RegisterConcrete(MsgUnbindProxy{}, "gridfxchain/staking/MsgUnbindProxy", nil)
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
