package ante

import (
	"fmt"
	ethcmn "github.com/ethereum/go-ethereum/common"
	ethermint "github.com/gridfx/fxchain/app/types"
	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/x/auth"
	evmtypes "github.com/gridfx/fxchain/x/evm/types"
	"github.com/gridfx/fxchain/x/gov/types"
	"github.com/gridfx/fxchain/x/params"
	paramstypes "github.com/gridfx/fxchain/x/params/types"
	stakingkeeper "github.com/gridfx/fxchain/x/staking"
	wasmtypes "github.com/gridfx/fxchain/x/wasm/types"
)

type AnteDecorator struct {
	sk stakingkeeper.Keeper
	ak auth.AccountKeeper
	pk params.Keeper
}

func NewAnteDecorator(k stakingkeeper.Keeper, ak auth.AccountKeeper, pk params.Keeper) AnteDecorator {
	return AnteDecorator{sk: k, ak: ak, pk: pk}
}

func (ad AnteDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	for _, m := range tx.GetMsgs() {
		switch msg := m.(type) {
		case types.MsgSubmitProposal:
			switch proposalType := msg.Content.(type) {
			case evmtypes.ManageContractByteCodeProposal:
				if !ad.sk.IsValidator(ctx, msg.Proposer) {
					return ctx, evmtypes.ErrCodeProposerMustBeValidator()
				}

				// check operation contract
				contract := ad.ak.GetAccount(ctx, proposalType.Contract)
				contractAcc, ok := contract.(*ethermint.EthAccount)
				if !ok || !contractAcc.IsContract() {
					return ctx, evmtypes.ErrNotContracAddress(fmt.Errorf(ethcmn.BytesToAddress(proposalType.Contract).String()))
				}

				//check substitute contract
				substitute := ad.ak.GetAccount(ctx, proposalType.SubstituteContract)
				substituteAcc, ok := substitute.(*ethermint.EthAccount)
				if !ok || !substituteAcc.IsContract() {
					return ctx, evmtypes.ErrNotContracAddress(fmt.Errorf(ethcmn.BytesToAddress(proposalType.SubstituteContract).String()))
				}
			case paramstypes.UpgradeProposal:
				if err := ad.pk.CheckMsgSubmitProposal(ctx, msg); err != nil {
					return ctx, err
				}
			case *wasmtypes.ExtraProposal:
				if !ad.sk.IsValidator(ctx, msg.Proposer) {
					return ctx, wasmtypes.ErrProposerMustBeValidator
				}
			}
		}
	}

	return next(ctx, tx, simulate)
}
