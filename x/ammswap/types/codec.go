package types

import (
	"github.com/gridfx/fxchain/libs/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgAddLiquidity{}, "okfxchain/ammswap/MsgAddLiquidity", nil)
	cdc.RegisterConcrete(MsgRemoveLiquidity{}, "okfxchain/ammswap/MsgRemoveLiquidity", nil)
	cdc.RegisterConcrete(MsgCreateExchange{}, "okfxchain/ammswap/MsgCreateExchange", nil)
	cdc.RegisterConcrete(MsgTokenToToken{}, "okfxchain/ammswap/MsgSwapToken", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
