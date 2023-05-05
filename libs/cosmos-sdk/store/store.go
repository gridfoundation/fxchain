package store

import (
	dbm "github.com/gridfx/fxchain/libs/tm-db"

	"github.com/gridfx/fxchain/libs/cosmos-sdk/store/cache"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/store/rootmulti"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
