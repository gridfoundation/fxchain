package types

import (
	"github.com/gridfx/fxchain/libs/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreatePool{}, "gridfxchain/farm/MsgCreatePool", nil)
	cdc.RegisterConcrete(MsgDestroyPool{}, "gridfxchain/farm/MsgDestroyPool", nil)
	cdc.RegisterConcrete(MsgLock{}, "gridfxchain/farm/MsgLock", nil)
	cdc.RegisterConcrete(MsgUnlock{}, "gridfxchain/farm/MsgUnlock", nil)
	cdc.RegisterConcrete(MsgClaim{}, "gridfxchain/farm/MsgClaim", nil)
	cdc.RegisterConcrete(MsgProvide{}, "gridfxchain/farm/MsgProvide", nil)
	cdc.RegisterConcrete(ManageWhiteListProposal{}, "gridfxchain/farm/ManageWhiteListProposal", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
