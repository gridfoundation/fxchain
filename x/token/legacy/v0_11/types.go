package v0_11

import (
	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"
	"github.com/gridfx/fxchain/x/token/legacy/v0_10"
)

const ModuleName = "token"

type (
	// all state that must be provided in genesis file
	GenesisState struct {
		Params       v0_10.Params     `json:"params"`
		Tokens       []Token          `json:"tokens"`
		LockedAssets []v0_10.AccCoins `json:"locked_assets"`
		LockedFees   []v0_10.AccCoins `json:"locked_fees"`
	}

	Token struct {
		Description         string         `json:"description" v2:"description"`                     // e.g. "FURY Group Global Utility Token"
		Symbol              string         `json:"symbol" v2:"symbol"`                               // e.g. "fury"
		OriginalSymbol      string         `json:"original_symbol" v2:"original_symbol"`             // e.g. "FURY"
		WholeName           string         `json:"whole_name" v2:"whole_name"`                       // e.g. "FURY"
		OriginalTotalSupply sdk.Dec        `json:"original_total_supply" v2:"original_total_supply"` // e.g. 1000000000.00000000
		Owner               sdk.AccAddress `json:"owner" v2:"owner"`                                 // e.g. did:fury:ex1rf9wr069pt64e58f2w3mjs9w72g8vemzg70a5e
		Mintable            bool           `json:"mintable" v2:"mintable"`                           // e.g. false
	}
)
