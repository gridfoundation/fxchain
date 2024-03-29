package types

import (
	codectypes "github.com/gridfx/fxchain/libs/cosmos-sdk/codec/types"
	"github.com/gridfx/fxchain/libs/ibc-go/modules/core/exported"
)

// RegisterInterfaces register the ibc interfaces submodule implementations to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*exported.ClientState)(nil),
		&ClientState{},
	)
}
