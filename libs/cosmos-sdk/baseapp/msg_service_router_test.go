package baseapp_test

import (
	okfxchaincodec "github.com/gridfx/fxchain/app/codec"
	"github.com/gridfx/fxchain/libs/ibc-go/testing/simapp"
	"github.com/gridfx/fxchain/x/evm"
	"os"
	"testing"

	"github.com/gridfx/fxchain/libs/tendermint/libs/log"

	dbm "github.com/gridfx/fxchain/libs/tm-db"
	"github.com/stretchr/testify/require"

	"github.com/gridfx/fxchain/libs/cosmos-sdk/baseapp"

	"github.com/gridfx/fxchain/x/evm/types/testdata"
)

func TestRegisterMsgService(t *testing.T) {
	db := dbm.NewMemDB()

	// Create an encoding config that doesn't register testdata Msg services.
	codecProxy, interfaceRegistry := okfxchaincodec.MakeCodecSuit(simapp.ModuleBasics)
	app := baseapp.NewBaseApp("test", log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, evm.TxDecoder(codecProxy))
	app.SetInterfaceRegistry(interfaceRegistry)
	require.Panics(t, func() {
		testdata.RegisterMsgServer(
			app.MsgServiceRouter(),
			testdata.MsgServerImpl{},
		)
	})

	// Register testdata Msg services, and rerun `RegisterService`.
	testdata.RegisterInterfaces(interfaceRegistry)
	require.NotPanics(t, func() {
		testdata.RegisterMsgServer(
			app.MsgServiceRouter(),
			testdata.MsgServerImpl{},
		)
	})
}

func TestRegisterMsgServiceTwice(t *testing.T) {
	// Setup baseapp.
	db := dbm.NewMemDB()
	codecProxy, interfaceRegistry := okfxchaincodec.MakeCodecSuit(simapp.ModuleBasics)
	app := baseapp.NewBaseApp("test", log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, evm.TxDecoder(codecProxy))
	app.SetInterfaceRegistry(interfaceRegistry)
	testdata.RegisterInterfaces(interfaceRegistry)

	// First time registering service shouldn't panic.
	require.NotPanics(t, func() {
		testdata.RegisterMsgServer(
			app.MsgServiceRouter(),
			testdata.MsgServerImpl{},
		)
	})

	// Second time should panic.
	require.Panics(t, func() {
		testdata.RegisterMsgServer(
			app.MsgServiceRouter(),
			testdata.MsgServerImpl{},
		)
	})
}
