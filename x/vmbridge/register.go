package vmbridge

import (
	"github.com/gridfx/fxchain/libs/cosmos-sdk/codec"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/types/module"
	"github.com/gridfx/fxchain/x/vmbridge/keeper"
	"github.com/gridfx/fxchain/x/wasm"
)

func RegisterServices(cfg module.Configurator, keeper keeper.Keeper) {
	RegisterMsgServer(cfg.MsgServer(), NewMsgServerImpl(keeper))
}

func GetWasmOpts(cdc *codec.ProtoCodec) wasm.Option {
	return wasm.WithMessageEncoders(RegisterSendToEvmEncoder(cdc))
}
