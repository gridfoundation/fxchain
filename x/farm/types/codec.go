package types

import (
	"github.com/gridfx/fxchain/libs/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreatePool{}, "fxchain/farm/MsgCreatePool", nil)
	cdc.RegisterConcrete(MsgDestroyPool{}, "fxchain/farm/MsgDestroyPool", nil)
	cdc.RegisterConcrete(MsgLock{}, "fxchain/farm/MsgLock", nil)
	cdc.RegisterConcrete(MsgUnlock{}, "fxchain/farm/MsgUnlock", nil)
	cdc.RegisterConcrete(MsgClaim{}, "fxchain/farm/MsgClaim", nil)
	cdc.RegisterConcrete(MsgProvide{}, "fxchain/farm/MsgProvide", nil)
	cdc.RegisterConcrete(ManageWhiteListProposal{}, "fxchain/farm/ManageWhiteListProposal", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
