package gov

import (
	"strings"
	"testing"

	abci "github.com/gridfx/fxchain/libs/tendermint/abci/types"

	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"

	"github.com/stretchr/testify/require"
)

func TestInvalidMsg(t *testing.T) {
	k := Keeper{}
	h := NewHandler(k)

	res, err := h(sdk.NewContext(nil, abci.Header{}, false, nil), sdk.NewTestMsg())
	require.Error(t, err)
	require.Nil(t, res)
	require.True(t, strings.Contains(err.Error(), "unrecognized gov message type"))
}
