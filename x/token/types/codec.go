package types

import (
	"github.com/gridfx/fxchain/libs/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgTokenIssue{}, "fxchain/token/MsgIssue", nil)
	cdc.RegisterConcrete(MsgTokenBurn{}, "fxchain/token/MsgBurn", nil)
	cdc.RegisterConcrete(MsgTokenMint{}, "fxchain/token/MsgMint", nil)
	cdc.RegisterConcrete(MsgMultiSend{}, "fxchain/token/MsgMultiTransfer", nil)
	cdc.RegisterConcrete(MsgSend{}, "fxchain/token/MsgTransfer", nil)
	cdc.RegisterConcrete(MsgTransferOwnership{}, "fxchain/token/MsgTransferOwnership", nil)
	cdc.RegisterConcrete(MsgConfirmOwnership{}, "fxchain/token/MsgConfirmOwnership", nil)
	cdc.RegisterConcrete(MsgTokenModify{}, "fxchain/token/MsgModify", nil)

	// for test
	//cdc.RegisterConcrete(MsgTokenDestroy{}, "fxchain/token/MsgDestroy", nil)
}

// generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
