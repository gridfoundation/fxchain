package types

import (
	"time"

	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"
)

// DefaultOwnershipConfirmWindow defines default confirm window
const DefaultOwnershipConfirmWindow = 24 * time.Hour

type ConfirmOwnership struct {
	Symbol  string         `json:"symbol"`
	Address sdk.AccAddress `json:"address"`
	Expire  time.Time      `json:expire`
}
