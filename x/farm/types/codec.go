package types

import (
	"github.com/gridfx/fxchain/libs/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreatePool{}, "okfxchain/farm/MsgCreatePool", nil)
	cdc.RegisterConcrete(MsgDestroyPool{}, "okfxchain/farm/MsgDestroyPool", nil)
	cdc.RegisterConcrete(MsgLock{}, "okfxchain/farm/MsgLock", nil)
	cdc.RegisterConcrete(MsgUnlock{}, "okfxchain/farm/MsgUnlock", nil)
	cdc.RegisterConcrete(MsgClaim{}, "okfxchain/farm/MsgClaim", nil)
	cdc.RegisterConcrete(MsgProvide{}, "okfxchain/farm/MsgProvide", nil)
	cdc.RegisterConcrete(ManageWhiteListProposal{}, "okfxchain/farm/ManageWhiteListProposal", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
