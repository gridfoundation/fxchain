package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gridfx/fxchain/app/logevents"
	"github.com/gridfx/fxchain/cmd/fxchaind/fss"
	"github.com/gridfx/fxchain/cmd/fxchaind/mpt"

	"github.com/gridfx/fxchain/app/rpc"
	evmtypes "github.com/gridfx/fxchain/x/evm/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	abci "github.com/gridfx/fxchain/libs/tendermint/abci/types"
	tmamino "github.com/gridfx/fxchain/libs/tendermint/crypto/encoding/amino"
	"github.com/gridfx/fxchain/libs/tendermint/crypto/multisig"
	"github.com/gridfx/fxchain/libs/tendermint/libs/cli"
	"github.com/gridfx/fxchain/libs/tendermint/libs/log"
	tmtypes "github.com/gridfx/fxchain/libs/tendermint/types"
	dbm "github.com/gridfx/fxchain/libs/tm-db"

	"github.com/gridfx/fxchain/libs/cosmos-sdk/baseapp"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/client/flags"
	clientkeys "github.com/gridfx/fxchain/libs/cosmos-sdk/client/keys"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/crypto/keys"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/server"
	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/x/auth"

	"github.com/gridfx/fxchain/app"
	"github.com/gridfx/fxchain/app/codec"
	"github.com/gridfx/fxchain/app/crypto/ethsecp256k1"
	fxchain "github.com/gridfx/fxchain/app/types"
	"github.com/gridfx/fxchain/cmd/client"
	"github.com/gridfx/fxchain/x/genutil"
	genutilcli "github.com/gridfx/fxchain/x/genutil/client/cli"
	genutiltypes "github.com/gridfx/fxchain/x/genutil/types"
	"github.com/gridfx/fxchain/x/staking"
)

const flagInvCheckPeriod = "inv-check-period"
const FxcEnvPrefix = "GRIDFXCHAIN"

var invCheckPeriod uint

func main() {
	cobra.EnableCommandSorting = false

	codecProxy, registry := codec.MakeCodecSuit(app.ModuleBasics)

	tmamino.RegisterKeyType(ethsecp256k1.PubKey{}, ethsecp256k1.PubKeyName)
	tmamino.RegisterKeyType(ethsecp256k1.PrivKey{}, ethsecp256k1.PrivKeyName)
	multisig.RegisterKeyType(ethsecp256k1.PubKey{}, ethsecp256k1.PubKeyName)

	keys.CryptoCdc = codecProxy.GetCdc()
	genutil.ModuleCdc = codecProxy.GetCdc()
	genutiltypes.ModuleCdc = codecProxy.GetCdc()
	clientkeys.KeysCdc = codecProxy.GetCdc()

	config := sdk.GetConfig()
	fxchain.SetBech32Prefixes(config)
	fxchain.SetBip44CoinType(config)
	config.Seal()

	ctx := server.NewDefaultContext()

	rootCmd := &cobra.Command{
		Use:               "fxchaind",
		Short:             "ExChain App Daemon (server)",
		PersistentPreRunE: preRun(ctx),
	}
	// CLI commands to initialize the chain
	rootCmd.AddCommand(
		client.ValidateChainID(
			genutilcli.InitCmd(ctx, codecProxy.GetCdc(), app.ModuleBasics, app.DefaultNodeHome),
		),
		genutilcli.CollectGenTxsCmd(ctx, codecProxy.GetCdc(), auth.GenesisAccountIterator{}, app.DefaultNodeHome),
		genutilcli.MigrateGenesisCmd(ctx, codecProxy.GetCdc()),
		genutilcli.GenTxCmd(
			ctx, codecProxy.GetCdc(), app.ModuleBasics, staking.AppModuleBasic{}, auth.GenesisAccountIterator{},
			app.DefaultNodeHome, app.DefaultCLIHome,
		),
		genutilcli.ValidateGenesisCmd(ctx, codecProxy.GetCdc(), app.ModuleBasics),
		client.TestnetCmd(ctx, codecProxy.GetCdc(), app.ModuleBasics, auth.GenesisAccountIterator{}),
		replayCmd(ctx, client.RegisterAppFlag, codecProxy, newApp, registry, registerRoutes),
		repairStateCmd(ctx),
		displayStateCmd(ctx),
		mpt.MptCmd(ctx),
		fss.Command(ctx),
		// AddGenesisAccountCmd allows users to add accounts to the genesis file
		AddGenesisAccountCmd(ctx, codecProxy.GetCdc(), app.DefaultNodeHome, app.DefaultCLIHome),
		flags.NewCompletionCmd(rootCmd, true),
		dataCmd(ctx),
		exportAppCmd(ctx),
		iaviewerCmd(ctx, codecProxy.GetCdc()),
		subscribeCmd(codecProxy.GetCdc()),
	)

	subFunc := func(logger log.Logger) log.Subscriber {
		return logevents.NewProvider(logger)
	}
	// Tendermint node base commands
	server.AddCommands(ctx, codecProxy, registry, rootCmd, newApp, closeApp, exportAppStateAndTMValidators,
		registerRoutes, client.RegisterAppFlag, app.PreRun, subFunc)

	// precheck flag syntax
	preCheckLongFlagSyntax()

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, FxcEnvPrefix, app.DefaultNodeHome)
	rootCmd.PersistentFlags().UintVar(&invCheckPeriod, flagInvCheckPeriod,
		0, "Assert registered invariants every N blocks")
	rootCmd.PersistentFlags().Bool(server.FlagGops, false, "Enable gops metrics collection")

	initEnv()
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func initEnv() {
	checkSetEnv("mempool_size", "200000")
	checkSetEnv("mempool_cache_size", "300000")
	checkSetEnv("mempool_force_recheck_gap", "2000")
	checkSetEnv("mempool_recheck", "false")
	checkSetEnv("consensus_timeout_commit", fmt.Sprintf("%dms", tmtypes.TimeoutCommit))
}

func checkSetEnv(envName string, value string) {
	realEnvName := FxcEnvPrefix + "_" + strings.ToUpper(envName)
	_, ok := os.LookupEnv(realEnvName)
	if !ok {
		_ = os.Setenv(realEnvName, value)
	}
}

func closeApp(iApp abci.Application) {
	fmt.Println("Close App")
	app := iApp.(*app.FXIronxChainApp)
	app.StopBaseApp()
	evmtypes.CloseIndexer()
	rpc.CloseEthBackend()
	app.EvmKeeper.Watcher.Stop()
}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	pruningOpts, err := server.GetPruningOptionsFromFlags()
	if err != nil {
		panic(err)
	}

	return app.NewFXIronxChainApp(
		logger,
		db,
		traceStore,
		true,
		map[int64]bool{},
		0,
		baseapp.SetPruning(pruningOpts),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
		baseapp.SetHaltHeight(uint64(viper.GetInt(server.FlagHaltHeight))),
	)
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
) (json.RawMessage, []tmtypes.GenesisValidator, error) {
	var ethermintApp *app.FXIronxChainApp
	if height != -1 {
		ethermintApp = app.NewFXIronxChainApp(logger, db, traceStore, false, map[int64]bool{}, 0)

		if err := ethermintApp.LoadHeight(height); err != nil {
			return nil, nil, err
		}
	} else {
		ethermintApp = app.NewFXIronxChainApp(logger, db, traceStore, true, map[int64]bool{}, 0)
	}

	return ethermintApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}

// All long flag must be in k=v format
func preCheckLongFlagSyntax() {
	params := os.Args[1:]
	for _, f := range params {
		tf := strings.TrimSpace(f)

		if strings.ToUpper(tf) == "TRUE" ||
			strings.ToUpper(tf) == "FALSE" {
			fmt.Fprintf(os.Stderr, "ERROR: Invalid parameter,"+
				" boolean flag should be --flag=true or --flag=false \n")
			os.Exit(1)
		}
	}
}

func preRun(ctx *server.Context) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		setReplayDefaultFlag()
		return server.PersistentPreRunEFn(ctx)(cmd, args)
	}
}
