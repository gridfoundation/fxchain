package types

import "github.com/gridfx/fxchain/libs/cosmos-sdk/codec"

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgList{}, "gridfxchain/dex/MsgList", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "gridfxchain/dex/MsgDeposit", nil)
	cdc.RegisterConcrete(MsgWithdraw{}, "gridfxchain/dex/MsgWithdraw", nil)
	cdc.RegisterConcrete(MsgTransferOwnership{}, "gridfxchain/dex/MsgTransferTradingPairOwnership", nil)
	cdc.RegisterConcrete(MsgConfirmOwnership{}, "gridfxchain/dex/MsgConfirmOwnership", nil)
	cdc.RegisterConcrete(DelistProposal{}, "gridfxchain/dex/DelistProposal", nil)
	cdc.RegisterConcrete(MsgCreateOperator{}, "gridfxchain/dex/CreateOperator", nil)
	cdc.RegisterConcrete(MsgUpdateOperator{}, "gridfxchain/dex/UpdateOperator", nil)
}

// ModuleCdc represents generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
