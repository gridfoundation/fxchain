package bank

import (
	"context"
	"github.com/gorilla/mux"
	clictx "github.com/gridfx/fxchain/libs/cosmos-sdk/client/context"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/codec"
	anytypes "github.com/gridfx/fxchain/libs/cosmos-sdk/codec/types"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/types/module"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/x/bank/internal/typesadapter"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
)

var (
	_ module.AppModuleBasicAdapter = AppModuleBasic{}
)

func (b AppModuleBasic) RegisterInterfaces(registry anytypes.InterfaceRegistry) {
	typesadapter.RegisterInterface(registry)
}

func (b AppModuleBasic) RegisterGRPCGatewayRoutes(cliContext clictx.CLIContext, serveMux *runtime.ServeMux) {
	typesadapter.RegisterQueryHandlerClient(context.Background(), serveMux, typesadapter.NewQueryClient(cliContext))
}

func (b AppModuleBasic) GetTxCmdV2(cdc *codec.CodecProxy, reg anytypes.InterfaceRegistry) *cobra.Command {
	return nil
}

func (b AppModuleBasic) GetQueryCmdV2(cdc *codec.CodecProxy, reg anytypes.InterfaceRegistry) *cobra.Command {
	return nil
}

func (b AppModuleBasic) RegisterRouterForGRPC(cliCtx clictx.CLIContext, r *mux.Router) {}
