package feesplit

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"
	sdkerrors "github.com/gridfx/fxchain/libs/cosmos-sdk/types/errors"
	"github.com/gridfx/fxchain/x/common"
	"github.com/gridfx/fxchain/x/feesplit/types"
	govTypes "github.com/gridfx/fxchain/x/gov/types"
)

// NewProposalHandler handles "gov" type message in "feesplit"
func NewProposalHandler(k *Keeper) govTypes.Handler {
	return func(ctx sdk.Context, proposal *govTypes.Proposal) (err sdk.Error) {
		switch content := proposal.Content.(type) {
		case types.FeeSplitSharesProposal:
			return handleFeeSplitSharesProposal(ctx, k, content)
		default:
			return common.ErrUnknownProposalType(types.DefaultCodespace, content.ProposalType())
		}
	}
}

func handleFeeSplitSharesProposal(ctx sdk.Context, k *Keeper, p types.FeeSplitSharesProposal) sdk.Error {
	for _, share := range p.Shares {
		contract := ethcommon.HexToAddress(share.ContractAddr)
		_, found := k.GetFeeSplit(ctx, contract)
		if !found {
			return sdkerrors.Wrapf(
				types.ErrFeeSplitContractNotRegistered,
				"contract %s is not registered", share.ContractAddr,
			)
		}

		k.SetContractShare(ctx, contract, share.Share)
	}
	return nil
}
