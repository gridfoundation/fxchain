package types

import sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"

// constants
const (
	// module name
	ModuleName = "feesplit"
	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName
	// RouterKey to be used for message routing
	RouterKey = ModuleName

	QueryParameters              = "params"
	QueryFeeSplits               = "fee-splits"
	QueryFeeSplit                = "fee-split"
	QueryDeployerFeeSplits       = "deployer-fee-splits"
	QueryDeployerFeeSplitsDetail = "deployer-fee-splits-detail"
	QueryWithdrawerFeeSplits     = "withdrawer-fee-splits"
)

// prefix bytes for the fees persistent store
const (
	prefixFeeSplit = iota + 1
	prefixDeployer
	prefixWithdrawer
	prefixContractShare
)

// KVStore key prefixes
var (
	KeyPrefixFeeSplit      = []byte{prefixFeeSplit}
	KeyPrefixDeployer      = []byte{prefixDeployer}
	KeyPrefixWithdrawer    = []byte{prefixWithdrawer}
	KeyPrefixContractShare = []byte{prefixContractShare}
)

// GetKeyPrefixDeployer returns the KVStore key prefix for storing
// registered fee split contract for a deployer
func GetKeyPrefixDeployer(deployerAddress sdk.AccAddress) []byte {
	return append(KeyPrefixDeployer, deployerAddress.Bytes()...)
}

// GetKeyPrefixWithdrawer returns the KVStore key prefix for storing
// registered fee split contract for a withdrawer
func GetKeyPrefixWithdrawer(withdrawerAddress sdk.AccAddress) []byte {
	return append(KeyPrefixWithdrawer, withdrawerAddress.Bytes()...)
}
