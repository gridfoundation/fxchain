package appstatus

import (
	"fmt"
	"math"
	"path/filepath"

	bam "github.com/gridfx/fxchain/libs/cosmos-sdk/baseapp"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/client/flags"
	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/x/auth"
	capabilitytypes "github.com/gridfx/fxchain/libs/cosmos-sdk/x/capability/types"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/x/mint"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/x/params"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/x/supply"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/x/upgrade"
	"github.com/gridfx/fxchain/libs/iavl"
	ibctransfertypes "github.com/gridfx/fxchain/libs/ibc-go/modules/apps/transfer/types"
	ibchost "github.com/gridfx/fxchain/libs/ibc-go/modules/core/24-host"
	dbm "github.com/gridfx/fxchain/libs/tm-db"
	"github.com/gridfx/fxchain/x/ammswap"
	dex "github.com/gridfx/fxchain/x/dex/types"
	distr "github.com/gridfx/fxchain/x/distribution"
	"github.com/gridfx/fxchain/x/erc20"
	"github.com/gridfx/fxchain/x/evidence"
	"github.com/gridfx/fxchain/x/evm"
	"github.com/gridfx/fxchain/x/farm"
	"github.com/gridfx/fxchain/x/feesplit"
	"github.com/gridfx/fxchain/x/gov"
	"github.com/gridfx/fxchain/x/order"
	"github.com/gridfx/fxchain/x/slashing"
	staking "github.com/gridfx/fxchain/x/staking/types"
	token "github.com/gridfx/fxchain/x/token/types"
	"github.com/spf13/viper"
)

const (
	applicationDB = "application"
	dbFolder      = "data"
)

func GetAllStoreKeys() []string {
	return []string{
		bam.MainStoreKey, auth.StoreKey, staking.StoreKey,
		supply.StoreKey, mint.StoreKey, distr.StoreKey, slashing.StoreKey,
		gov.StoreKey, params.StoreKey, upgrade.StoreKey, evidence.StoreKey,
		evm.StoreKey, token.StoreKey, token.KeyLock, dex.StoreKey, dex.TokenPairStoreKey,
		order.OrderStoreKey, ammswap.StoreKey, farm.StoreKey, ibctransfertypes.StoreKey, capabilitytypes.StoreKey,
		ibchost.StoreKey,
		erc20.StoreKey,
		// mpt.StoreKey,
		// wasm.StoreKey,
		feesplit.StoreKey,
	}
}

func IsFastStorageStrategy() bool {
	return checkFastStorageStrategy(GetAllStoreKeys())
}

func checkFastStorageStrategy(storeKeys []string) bool {
	home := viper.GetString(flags.FlagHome)
	dataDir := filepath.Join(home, dbFolder)
	db, err := sdk.NewDB(applicationDB, dataDir)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for _, v := range storeKeys {
		if !isFss(db, v) {
			return false
		}
	}

	return true
}

func isFss(db dbm.DB, storeKey string) bool {
	prefix := fmt.Sprintf("s/k:%s/", storeKey)
	prefixDB := dbm.NewPrefixDB(db, []byte(prefix))

	return iavl.IsFastStorageStrategy(prefixDB)
}

func GetFastStorageVersion() int64 {
	home := viper.GetString(flags.FlagHome)
	dataDir := filepath.Join(home, dbFolder)
	db, err := sdk.NewDB(applicationDB, dataDir)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	storeKeys := GetAllStoreKeys()
	var ret int64 = math.MaxInt64
	for _, v := range storeKeys {
		version := getVersion(db, v)
		if version < ret {
			ret = version
		}
	}

	return ret
}

func getVersion(db dbm.DB, storeKey string) int64 {
	prefix := fmt.Sprintf("s/k:%s/", storeKey)
	prefixDB := dbm.NewPrefixDB(db, []byte(prefix))

	version, _ := iavl.GetFastStorageVersion(prefixDB)

	return version
}
