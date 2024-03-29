package simulation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	tmkv "github.com/gridfx/fxchain/libs/tendermint/libs/kv"

	"github.com/gridfx/fxchain/libs/cosmos-sdk/codec"
	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/x/mint/internal/types"
)

func makeTestCodec() (cdc *codec.Codec) {
	cdc = codec.New()
	sdk.RegisterCodec(cdc)
	return
}

func TestDecodeStore(t *testing.T) {
	cdc := makeTestCodec()
	minter := types.NewMinter(sdk.OneDec(), sdk.NewDec(15))

	kvPairs := tmkv.Pairs{
		tmkv.Pair{Key: types.MinterKey, Value: cdc.MustMarshalBinaryLengthPrefixed(minter)},
		tmkv.Pair{Key: []byte{0x99}, Value: []byte{0x99}},
	}
	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Minter", fmt.Sprintf("%v\n%v", minter, minter)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { DecodeStore(cdc, kvPairs[i], kvPairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, DecodeStore(cdc, kvPairs[i], kvPairs[i]), tt.name)
			}
		})
	}
}
