package types

import (
	"github.com/gridfx/fxchain/libs/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgTokenIssue{}, "okfxchain/token/MsgIssue", nil)
	cdc.RegisterConcrete(MsgTokenBurn{}, "okfxchain/token/MsgBurn", nil)
	cdc.RegisterConcrete(MsgTokenMint{}, "okfxchain/token/MsgMint", nil)
	cdc.RegisterConcrete(MsgMultiSend{}, "okfxchain/token/MsgMultiTransfer", nil)
	cdc.RegisterConcrete(MsgSend{}, "okfxchain/token/MsgTransfer", nil)
	cdc.RegisterConcrete(MsgTransferOwnership{}, "okfxchain/token/MsgTransferOwnership", nil)
	cdc.RegisterConcrete(MsgConfirmOwnership{}, "okfxchain/token/MsgConfirmOwnership", nil)
	cdc.RegisterConcrete(MsgTokenModify{}, "okfxchain/token/MsgModify", nil)

	// for test
	//cdc.RegisterConcrete(MsgTokenDestroy{}, "okfxchain/token/MsgDestroy", nil)
}

// generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
