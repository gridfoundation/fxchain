package types

import "github.com/gridfx/fxchain/libs/cosmos-sdk/codec"

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgList{}, "okfxchain/dex/MsgList", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "okfxchain/dex/MsgDeposit", nil)
	cdc.RegisterConcrete(MsgWithdraw{}, "okfxchain/dex/MsgWithdraw", nil)
	cdc.RegisterConcrete(MsgTransferOwnership{}, "okfxchain/dex/MsgTransferTradingPairOwnership", nil)
	cdc.RegisterConcrete(MsgConfirmOwnership{}, "okfxchain/dex/MsgConfirmOwnership", nil)
	cdc.RegisterConcrete(DelistProposal{}, "okfxchain/dex/DelistProposal", nil)
	cdc.RegisterConcrete(MsgCreateOperator{}, "okfxchain/dex/CreateOperator", nil)
	cdc.RegisterConcrete(MsgUpdateOperator{}, "okfxchain/dex/UpdateOperator", nil)
}

// ModuleCdc represents generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
