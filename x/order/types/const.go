package types

import (
	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"
	"github.com/gridfx/fxchain/x/common"
)

// nolint
const (
	FeeTypeOrderNew     = "new"
	FeeTypeOrderCancel  = "cancel"
	FeeTypeOrderExpire  = "expire"
	FeeTypeOrderDeal    = "deal"
	FeeTypeOrderReceive = "receive"
	TestTokenPair       = common.TestToken + "_" + sdk.DefaultBondDenom
	BuyOrder            = "BUY"
	SellOrder           = "SELL"
)
