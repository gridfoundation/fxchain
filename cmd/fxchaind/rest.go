package main

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/gridfx/fxchain/app"
	"github.com/gridfx/fxchain/app/rpc"
	"github.com/gridfx/fxchain/app/types"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/client"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/client/lcd"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/server"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/types/tx"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/x/auth"
	authrest "github.com/gridfx/fxchain/libs/cosmos-sdk/x/auth/client/rest"
	bankrest "github.com/gridfx/fxchain/libs/cosmos-sdk/x/bank/client/rest"
	mintclient "github.com/gridfx/fxchain/libs/cosmos-sdk/x/mint/client"
	mintrest "github.com/gridfx/fxchain/libs/cosmos-sdk/x/mint/client/rest"
	supplyrest "github.com/gridfx/fxchain/libs/cosmos-sdk/x/supply/client/rest"
	ibctransferrest "github.com/gridfx/fxchain/libs/ibc-go/modules/apps/transfer/client/rest"
	ammswaprest "github.com/gridfx/fxchain/x/ammswap/client/rest"
	dexclient "github.com/gridfx/fxchain/x/dex/client"
	dexrest "github.com/gridfx/fxchain/x/dex/client/rest"
	dist "github.com/gridfx/fxchain/x/distribution"
	distr "github.com/gridfx/fxchain/x/distribution"
	distrest "github.com/gridfx/fxchain/x/distribution/client/rest"
	erc20client "github.com/gridfx/fxchain/x/erc20/client"
	erc20rest "github.com/gridfx/fxchain/x/erc20/client/rest"
	evmclient "github.com/gridfx/fxchain/x/evm/client"
	evmrest "github.com/gridfx/fxchain/x/evm/client/rest"
	farmclient "github.com/gridfx/fxchain/x/farm/client"
	farmrest "github.com/gridfx/fxchain/x/farm/client/rest"
	fsrest "github.com/gridfx/fxchain/x/feesplit/client/rest"
	govrest "github.com/gridfx/fxchain/x/gov/client/rest"
	orderrest "github.com/gridfx/fxchain/x/order/client/rest"
	paramsclient "github.com/gridfx/fxchain/x/params/client"
	slashingrest "github.com/gridfx/fxchain/x/slashing/client/rest"
	stakingrest "github.com/gridfx/fxchain/x/staking/client/rest"
	"github.com/gridfx/fxchain/x/token"
	tokensrest "github.com/gridfx/fxchain/x/token/client/rest"
	wasmrest "github.com/gridfx/fxchain/x/wasm/client/rest"
	"github.com/gridfx/fxchain/x/wasm/proxy"
)

// registerRoutes registers the routes from the different modules for the LCD.
// NOTE: details on the routes added for each module are in the module documentation
// NOTE: If making updates here you also need to update the test helper in client/lcd/test_helper.go
func registerRoutes(rs *lcd.RestServer) {
	registerGrpc(rs)
	rpc.RegisterRoutes(rs)
	pathPrefix := viper.GetString(server.FlagRestPathPrefix)
	if pathPrefix == "" {
		pathPrefix = types.EthBech32Prefix
	}
	registerRoutesV1(rs, pathPrefix)
	registerRoutesV2(rs, pathPrefix)
	proxy.SetCliContext(rs.CliCtx)
}

func registerGrpc(rs *lcd.RestServer) {
	app.ModuleBasics.RegisterGRPCGatewayRoutes(rs.CliCtx, rs.GRPCGatewayRouter)
	app.ModuleBasics.RegisterRPCRouterForGRPC(rs.CliCtx, rs.Mux)
	tx.RegisterGRPCGatewayRoutes(rs.CliCtx, rs.GRPCGatewayRouter)
}

func registerRoutesV1(rs *lcd.RestServer, pathPrefix string) {
	v1Router := rs.Mux.PathPrefix(fmt.Sprintf("/%s/v1", pathPrefix)).Name("v1").Subrouter()
	client.RegisterRoutes(rs.CliCtx, v1Router)
	authrest.RegisterRoutes(rs.CliCtx, v1Router, auth.StoreKey)
	bankrest.RegisterRoutes(rs.CliCtx, v1Router)
	stakingrest.RegisterRoutes(rs.CliCtx, v1Router)
	slashingrest.RegisterRoutes(rs.CliCtx, v1Router)
	distrest.RegisterRoutes(rs.CliCtx, v1Router, dist.StoreKey)

	orderrest.RegisterRoutes(rs.CliCtx, v1Router)
	tokensrest.RegisterRoutes(rs.CliCtx, v1Router, token.StoreKey)
	dexrest.RegisterRoutes(rs.CliCtx, v1Router)
	ammswaprest.RegisterRoutes(rs.CliCtx, v1Router)
	supplyrest.RegisterRoutes(rs.CliCtx, v1Router)
	farmrest.RegisterRoutes(rs.CliCtx, v1Router)
	evmrest.RegisterRoutes(rs.CliCtx, v1Router)
	erc20rest.RegisterRoutes(rs.CliCtx, v1Router)
	wasmrest.RegisterRoutes(rs.CliCtx, v1Router)
	fsrest.RegisterRoutes(rs.CliCtx, v1Router)
	govrest.RegisterRoutes(rs.CliCtx, v1Router,
		[]govrest.ProposalRESTHandler{
			paramsclient.ProposalHandler.RESTHandler(rs.CliCtx),
			distr.CommunityPoolSpendProposalHandler.RESTHandler(rs.CliCtx),
			distr.ChangeDistributionTypeProposalHandler.RESTHandler(rs.CliCtx),
			distr.WithdrawRewardEnabledProposalHandler.RESTHandler(rs.CliCtx),
			distr.RewardTruncatePrecisionProposalHandler.RESTHandler(rs.CliCtx),
			dexclient.DelistProposalHandler.RESTHandler(rs.CliCtx),
			farmclient.ManageWhiteListProposalHandler.RESTHandler(rs.CliCtx),
			evmclient.ManageContractDeploymentWhitelistProposalHandler.RESTHandler(rs.CliCtx),
			evmclient.ManageSysContractAddressProposalHandler.RESTHandler(rs.CliCtx),
			evmclient.ManageContractByteCodeProposalHandler.RESTHandler(rs.CliCtx),
			mintclient.ManageTreasuresProposalHandler.RESTHandler(rs.CliCtx),
			mintclient.ModifyNextBlockUpdateProposalHandler.RESTHandler(rs.CliCtx),
			erc20client.TokenMappingProposalHandler.RESTHandler(rs.CliCtx),
		},
	)
	mintrest.RegisterRoutes(rs.CliCtx, v1Router)
	ibctransferrest.RegisterOriginRPCRoutersForGRPC(rs.CliCtx, v1Router)
}

func registerRoutesV2(rs *lcd.RestServer, pathPrefix string) {
	v2Router := rs.Mux.PathPrefix(fmt.Sprintf("/%s/v2", pathPrefix)).Name("v1").Subrouter()
	client.RegisterRoutes(rs.CliCtx, v2Router)
	authrest.RegisterRoutes(rs.CliCtx, v2Router, auth.StoreKey)
	bankrest.RegisterRoutes(rs.CliCtx, v2Router)
	stakingrest.RegisterRoutes(rs.CliCtx, v2Router)
	distrest.RegisterRoutes(rs.CliCtx, v2Router, dist.StoreKey)
	orderrest.RegisterRoutesV2(rs.CliCtx, v2Router)
	tokensrest.RegisterRoutesV2(rs.CliCtx, v2Router, token.StoreKey)
	fsrest.RegisterRoutesV2(rs.CliCtx, v2Router)
}
