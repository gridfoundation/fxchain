package types

import "github.com/gridfx/fxchain/libs/cosmos-sdk/codec"

var ModuleCdc *codec.Codec

func init(){
	ModuleCdc = codec.New()
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
