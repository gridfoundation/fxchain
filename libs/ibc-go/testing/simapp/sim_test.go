package simapp

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"testing"

	abci "github.com/gridfx/fxchain/libs/tendermint/abci/types"
	"github.com/gridfx/fxchain/libs/tendermint/libs/log"
	dbm "github.com/gridfx/fxchain/libs/tm-db"
	"github.com/stretchr/testify/require"

	"github.com/gridfx/fxchain/libs/cosmos-sdk/baseapp"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/store"
	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"
	authtypes "github.com/gridfx/fxchain/libs/cosmos-sdk/x/auth/types"
	banktypes "github.com/gridfx/fxchain/libs/cosmos-sdk/x/bank"
	capabilitytypes "github.com/gridfx/fxchain/libs/cosmos-sdk/x/capability/types"
	distrtypes "github.com/gridfx/fxchain/libs/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/gridfx/fxchain/libs/cosmos-sdk/x/evidence"
	govtypes "github.com/gridfx/fxchain/libs/cosmos-sdk/x/gov/types"
	minttypes "github.com/gridfx/fxchain/libs/cosmos-sdk/x/mint"
	paramtypes "github.com/gridfx/fxchain/libs/cosmos-sdk/x/params/types"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/x/simulation"
	slashingtypes "github.com/gridfx/fxchain/libs/cosmos-sdk/x/slashing"
	stakingtypes "github.com/gridfx/fxchain/libs/cosmos-sdk/x/staking/types"
	ibctransfertypes "github.com/gridfx/fxchain/libs/ibc-go/modules/apps/transfer/types"
	ibchost "github.com/gridfx/fxchain/libs/ibc-go/modules/core/24-host"
	"github.com/gridfx/fxchain/libs/ibc-go/testing/simapp/helpers"
)

// Get flags every time the simulator is run
func init() {
	GetSimulatorFlags()
}

type StoreKeysPrefixes struct {
	A        sdk.StoreKey
	B        sdk.StoreKey
	Prefixes [][]byte
}

// fauxMerkleModeOpt returns a BaseApp option to use a dbStoreAdapter instead of
// an IAVLStore for faster simulation speed.
func fauxMerkleModeOpt(bapp *baseapp.BaseApp) {
	bapp.SetFauxMerkleMode()
}

// interBlockCacheOpt returns a BaseApp option function that sets the persistent
// inter-block write-through cache.
func interBlockCacheOpt() func(*baseapp.BaseApp) {
	return baseapp.SetInterBlockCache(store.NewCommitKVStoreCacheManager())
}

func TestFullAppSimulation(t *testing.T) {
	config, db, dir, logger, skip, err := SetupSimulation("leveldb-app-sim", "Simulation")
	if skip {
		t.Skip("skipping application simulation")
	}
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		db.Close()
		require.NoError(t, os.RemoveAll(dir))
	}()

	app := NewSimApp(logger, db, nil, true, map[int64]bool{}, FlagPeriodValue, fauxMerkleModeOpt)
	require.Equal(t, "SimApp", app.Name())

	// run randomized simulation
	_, simParams, simErr := simulation.SimulateFromSeed(
		t,
		os.Stdout,
		app.BaseApp,
		AppStateFn(*app.AppCodec(), app.SimulationManager()),
		//		simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		SimulationOperations(app, *app.AppCodec().GetCdc(), config),
		app.ModuleAccountAddrs(),
		config,
	)

	// export state and simParams before the simulation error is checked
	err = CheckExportSimulation(app, config, simParams)
	require.NoError(t, err)
	require.NoError(t, simErr)

	if config.Commit {
		PrintStats(db)
	}
}

func TestAppImportExport(t *testing.T) {
	config, db, dir, logger, skip, err := SetupSimulation("leveldb-app-sim", "Simulation")
	if skip {
		t.Skip("skipping application import/export simulation")
	}
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		db.Close()
		require.NoError(t, os.RemoveAll(dir))
	}()

	app := NewSimApp(logger, db, nil, true, map[int64]bool{}, FlagPeriodValue, fauxMerkleModeOpt)
	require.Equal(t, "SimApp", app.Name())

	// Run randomized simulation
	_, simParams, simErr := simulation.SimulateFromSeed(
		t,
		os.Stdout,
		app.BaseApp,
		AppStateFn(*app.AppCodec(), app.SimulationManager()),
		//		simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		SimulationOperations(app, *app.AppCodec().GetCdc(), config),
		app.ModuleAccountAddrs(),
		config,
	)

	// export state and simParams before the simulation error is checked
	err = CheckExportSimulation(app, config, simParams)
	require.NoError(t, err)
	require.NoError(t, simErr)

	if config.Commit {
		PrintStats(db)
	}

	fmt.Printf("exporting genesis...\n")

	exported, _, err := app.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err)

	fmt.Printf("importing genesis...\n")

	_, newDB, newDir, _, _, err := SetupSimulation("leveldb-app-sim-2", "Simulation-2")
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		newDB.Close()
		require.NoError(t, os.RemoveAll(newDir))
	}()

	newApp := NewSimApp(log.NewNopLogger(), newDB, nil, true, map[int64]bool{}, FlagPeriodValue /*EmptyAppOptions{},*/, fauxMerkleModeOpt)
	require.Equal(t, "SimApp", newApp.Name())

	var genesisState GenesisState
	err = json.Unmarshal(exported, &genesisState)
	require.NoError(t, err)

	ctxA := app.NewContext(true, abci.Header{Height: app.LastBlockHeight()})
	ctxB := newApp.NewContext(true, abci.Header{Height: app.LastBlockHeight()})
	newApp.mm.InitGenesis(ctxB, genesisState)

	fmt.Printf("comparing stores...\n")

	storeKeysPrefixes := []StoreKeysPrefixes{
		{app.keys[authtypes.StoreKey], newApp.keys[authtypes.StoreKey], [][]byte{}},
		{app.keys[stakingtypes.StoreKey], newApp.keys[stakingtypes.StoreKey],
			[][]byte{
				stakingtypes.UnbondingQueueKey, stakingtypes.RedelegationQueueKey, stakingtypes.ValidatorQueueKey,
				stakingtypes.HistoricalInfoKey,
			}}, // ordering may change but it doesn't matter
		{app.keys[slashingtypes.StoreKey], newApp.keys[slashingtypes.StoreKey], [][]byte{}},
		{app.keys[minttypes.StoreKey], newApp.keys[minttypes.StoreKey], [][]byte{}},
		{app.keys[distrtypes.StoreKey], newApp.keys[distrtypes.StoreKey], [][]byte{}},
		// todo check bank types
		//		{app.keys[banktypes.StoreKey], newApp.keys[banktypes.StoreKey], [][]byte{banktypes.BalancesPrefix}},
		{app.keys[banktypes.ModuleName], newApp.keys[banktypes.ModuleName], [][]byte{}},
		//		{app.keys[paramtypes.StoreKey], newApp.keys[paramtypes.StoreKey], [][]byte{}},
		{app.keys[paramtypes.ModuleName], newApp.keys[paramtypes.ModuleName], [][]byte{}},
		{app.keys[govtypes.StoreKey], newApp.keys[govtypes.StoreKey], [][]byte{}},
		{app.keys[evidencetypes.StoreKey], newApp.keys[evidencetypes.StoreKey], [][]byte{}},
		{app.keys[capabilitytypes.StoreKey], newApp.keys[capabilitytypes.StoreKey], [][]byte{}},
		{app.keys[ibchost.StoreKey], newApp.keys[ibchost.StoreKey], [][]byte{}},
		{app.keys[ibctransfertypes.StoreKey], newApp.keys[ibctransfertypes.StoreKey], [][]byte{}},
		//{app.keys[authzkeeper.StoreKey], newApp.keys[authzkeeper.StoreKey], [][]byte{}},
	}

	for _, skp := range storeKeysPrefixes {
		storeA := ctxA.KVStore(skp.A)
		storeB := ctxB.KVStore(skp.B)

		failedKVAs, failedKVBs := sdk.DiffKVStores(storeA, storeB, skp.Prefixes)
		require.Equal(t, len(failedKVAs), len(failedKVBs), "unequal sets of key-values to compare")

		fmt.Printf("compared %d different key/value pairs between %s and %s\n", len(failedKVAs), skp.A, skp.B)
		require.Equal(t, len(failedKVAs), 0, GetSimulationLog(skp.A.Name(), app.SimulationManager().StoreDecoders, failedKVAs, failedKVBs))
	}
}

func TestAppSimulationAfterImport(t *testing.T) {
	config, db, dir, logger, skip, err := SetupSimulation("leveldb-app-sim", "Simulation")
	if skip {
		t.Skip("skipping application simulation after import")
	}
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		db.Close()
		require.NoError(t, os.RemoveAll(dir))
	}()

	app := NewSimApp(logger, db, nil, true, map[int64]bool{}, FlagPeriodValue, fauxMerkleModeOpt)
	require.Equal(t, "SimApp", app.Name())

	// Run randomized simulation
	stopEarly, simParams, simErr := simulation.SimulateFromSeed(
		t,
		os.Stdout,
		app.BaseApp,
		AppStateFn(*app.AppCodec(), app.SimulationManager()),
		//		simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		SimulationOperations(app, *app.AppCodec().GetCdc(), config),
		app.ModuleAccountAddrs(),
		config,
	)

	// export state and simParams before the simulation error is checked
	err = CheckExportSimulation(app, config, simParams)
	require.NoError(t, err)
	require.NoError(t, simErr)

	if config.Commit {
		PrintStats(db)
	}

	if stopEarly {
		fmt.Println("can't export or import a zero-validator genesis, exiting test...")
		return
	}

	fmt.Printf("exporting genesis...\n")

	exported, _, err := app.ExportAppStateAndValidators(true, []string{})
	require.NoError(t, err)

	fmt.Printf("importing genesis...\n")

	_, newDB, newDir, _, _, err := SetupSimulation("leveldb-app-sim-2", "Simulation-2")
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		newDB.Close()
		require.NoError(t, os.RemoveAll(newDir))
	}()

	newApp := NewSimApp(log.NewNopLogger(), newDB, nil, true, map[int64]bool{}, FlagPeriodValue, fauxMerkleModeOpt)

	require.Equal(t, "SimApp", newApp.Name())

	newApp.InitChain(abci.RequestInitChain{
		AppStateBytes: exported,
	})

	_, _, err = simulation.SimulateFromSeed(
		t,
		os.Stdout,
		newApp.GetBaseApp(),
		AppStateFn(*app.AppCodec(), app.SimulationManager()),
		//		simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		SimulationOperations(newApp, *newApp.AppCodec().GetCdc(), config),
		app.ModuleAccountAddrs(),
		config,
	)
	require.NoError(t, err)
}

// TODO: Make another test for the fuzzer itself, which just has noOp txs
// and doesn't depend on the application.
func TestAppStateDeterminism(t *testing.T) {
	if !FlagEnabledValue {
		t.Skip("skipping application simulation")
	}

	config := NewConfigFromFlags()
	config.InitialBlockHeight = 1
	config.ExportParamsPath = ""
	config.OnOperation = false
	config.AllInvariants = false
	config.ChainID = helpers.SimAppChainID

	numSeeds := 3
	numTimesToRunPerSeed := 5
	appHashList := make([]json.RawMessage, numTimesToRunPerSeed)

	for i := 0; i < numSeeds; i++ {
		config.Seed = rand.Int63()

		for j := 0; j < numTimesToRunPerSeed; j++ {
			var logger log.Logger
			if FlagVerboseValue {
				logger = log.TestingLogger()
			} else {
				logger = log.NewNopLogger()
			}

			db := dbm.NewMemDB()
			// todo MakeTestEncodingConfig EmptyAppOptions
			//	app := NewSimApp(logger, db, nil, true, map[int64]bool{}, DefaultNodeHome, FlagPeriodValue, MakeTestEncodingConfig(), EmptyAppOptions{}, interBlockCacheOpt())
			app := NewSimApp(logger, db, nil, true, map[int64]bool{}, FlagPeriodValue, interBlockCacheOpt())

			fmt.Printf(
				"running non-determinism simulation; seed %d: %d/%d, attempt: %d/%d\n",
				config.Seed, i+1, numSeeds, j+1, numTimesToRunPerSeed,
			)

			_, _, err := simulation.SimulateFromSeed(
				t,
				os.Stdout,
				app.BaseApp,
				AppStateFn(*app.AppCodec(), app.SimulationManager()),
				//		simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
				SimulationOperations(app, *app.AppCodec().GetCdc(), config),
				app.ModuleAccountAddrs(),
				config,
			)
			require.NoError(t, err)

			if config.Commit {
				PrintStats(db)
			}

			appHash := app.LastCommitID().Hash
			appHashList[j] = appHash

			if j != 0 {
				require.Equal(
					t, string(appHashList[0]), string(appHashList[j]),
					"non-determinism in seed %d: %d/%d, attempt: %d/%d\n", config.Seed, i+1, numSeeds, j+1, numTimesToRunPerSeed,
				)
			}
		}
	}
}
