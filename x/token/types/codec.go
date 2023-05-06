package types

import (
	"github.com/gridfx/fxchain/libs/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgTokenIssue{}, "gridfxchain/token/MsgIssue", nil)
	cdc.RegisterConcrete(MsgTokenBurn{}, "gridfxchain/token/MsgBurn", nil)
	cdc.RegisterConcrete(MsgTokenMint{}, "gridfxchain/token/MsgMint", nil)
	cdc.RegisterConcrete(MsgMultiSend{}, "gridfxchain/token/MsgMultiTransfer", nil)
	cdc.RegisterConcrete(MsgSend{}, "gridfxchain/token/MsgTransfer", nil)
	cdc.RegisterConcrete(MsgTransferOwnership{}, "gridfxchain/token/MsgTransferOwnership", nil)
	cdc.RegisterConcrete(MsgConfirmOwnership{}, "gridfxchain/token/MsgConfirmOwnership", nil)
	cdc.RegisterConcrete(MsgTokenModify{}, "gridfxchain/token/MsgModify", nil)

	// for test
	//cdc.RegisterConcrete(MsgTokenDestroy{}, "gridfxchain/token/MsgDestroy", nil)
}

// generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
