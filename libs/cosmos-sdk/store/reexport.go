package store

import (
	"github.com/gridfx/fxchain/libs/cosmos-sdk/store/types"
)

// Import cosmos-sdk/types/store.go for convenience.
// nolint
type (
	PruningOptions   = types.PruningOptions
	Store            = types.Store
	Committer        = types.Committer
	CommitStore      = types.CommitStore
	MultiStore       = types.MultiStore
	CacheMultiStore  = types.CacheMultiStore
	CommitMultiStore = types.CommitMultiStore
	KVStore          = types.KVStore
	KVPair           = types.KVPair
	Iterator         = types.Iterator
	CacheKVStore     = types.CacheKVStore
	CommitKVStore    = types.CommitKVStore
	CacheWrapper     = types.CacheWrapper
	CacheWrap        = types.CacheWrap
	CommitID         = types.CommitID
	StoreKey         = types.StoreKey
	StoreType        = types.StoreType
	Queryable        = types.Queryable
	TraceContext     = types.TraceContext
	Gas              = types.Gas
	GasMeter         = types.GasMeter
	GasConfig        = types.GasConfig
)
