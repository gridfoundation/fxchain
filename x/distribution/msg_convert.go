package distribution

import (
	"encoding/json"
	"errors"

	"github.com/gridfx/fxchain/libs/cosmos-sdk/baseapp"
	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"
	tmtypes "github.com/gridfx/fxchain/libs/tendermint/types"
	"github.com/gridfx/fxchain/x/common"
	"github.com/gridfx/fxchain/x/distribution/types"
)

var (
	ErrCheckSignerFail = errors.New("check signer fail")
)

func init() {
	RegisterConvert()
}

func RegisterConvert() {
	enableHeight := tmtypes.GetVenus3Height()
	baseapp.RegisterCmHandle("gridfxchain/distribution/MsgWithdrawDelegatorAllRewards", baseapp.NewCMHandle(ConvertWithdrawDelegatorAllRewardsMsg, enableHeight))
}

func ConvertWithdrawDelegatorAllRewardsMsg(data []byte, signers []sdk.AccAddress) (sdk.Msg, error) {
	newMsg := types.MsgWithdrawDelegatorAllRewards{}
	err := json.Unmarshal(data, &newMsg)
	if err != nil {
		return nil, err
	}
	err = newMsg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	if ok := common.CheckSignerAddress(signers, newMsg.GetSigners()); !ok {
		return nil, ErrCheckSignerFail
	}
	return newMsg, nil
}
