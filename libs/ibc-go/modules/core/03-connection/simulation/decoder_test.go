package simulation_test

import (
	"fmt"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/types/kv"
	tmkv "github.com/gridfx/fxchain/libs/tendermint/libs/kv"
	"testing"

	"github.com/stretchr/testify/require"

	// "github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/gridfx/fxchain/libs/ibc-go/modules/core/03-connection/simulation"
	"github.com/gridfx/fxchain/libs/ibc-go/modules/core/03-connection/types"
	host "github.com/gridfx/fxchain/libs/ibc-go/modules/core/24-host"
	"github.com/gridfx/fxchain/libs/ibc-go/testing/simapp"
)

func TestDecodeStore(t *testing.T) {
	app := simapp.Setup(false)
	cdc := app.AppCodec()

	connectionID := "connectionidone"

	connection := types.ConnectionEnd{
		ClientId: "clientidone",
		Versions: types.ExportedVersionsToProto(types.GetCompatibleVersions()),
	}

	paths := types.ClientPaths{
		Paths: []string{connectionID},
	}

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{
				Key: host.ClientConnectionsKey(connection.ClientId),
				//Value: cdc.MustMarshal(&paths),
				Value: cdc.GetProtocMarshal().MustMarshalBinaryBare(&paths),
			},
			{
				Key: host.ConnectionKey(connectionID),
				//Value: cdc.MustMarshal(&connection),
				Value: cdc.GetProtocMarshal().MustMarshalBinaryBare(&connection),
			},
			{
				Key:   []byte{0x99},
				Value: []byte{0x99},
			},
		},
	}
	tests := []struct {
		name        string
		expectedLog string
	}{
		{"ClientPaths", fmt.Sprintf("ClientPaths A: %v\nClientPaths B: %v", paths, paths)},
		{"ConnectionEnd", fmt.Sprintf("ConnectionEnd A: %v\nConnectionEnd B: %v", connection, connection)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			// 	res, found := simulation.NewDecodeStore(cdc, kvPairs.Pairs[i], kvPairs.Pairs[i])
			kvA := tmkv.Pair{
				Key:   kvPairs.Pairs[i].GetKey(),
				Value: kvPairs.Pairs[i].GetValue(),
			}
			res, found := simulation.NewDecodeStore(cdc, kvA, kvA)
			if i == len(tests)-1 {
				require.False(t, found, string(kvPairs.Pairs[i].Key))
				require.Empty(t, res, string(kvPairs.Pairs[i].Key))
			} else {
				require.True(t, found, string(kvPairs.Pairs[i].Key))
				require.Equal(t, tt.expectedLog, res, string(kvPairs.Pairs[i].Key))
			}
		})
	}
}
