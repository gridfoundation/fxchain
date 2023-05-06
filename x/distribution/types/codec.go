package types

import (
	"github.com/gridfx/fxchain/libs/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgWithdrawValidatorCommission{}, "gridfxchain/distribution/MsgWithdrawReward", nil)
	cdc.RegisterConcrete(MsgWithdrawDelegatorReward{}, "gridfxchain/distribution/MsgWithdrawDelegatorReward", nil)
	cdc.RegisterConcrete(MsgSetWithdrawAddress{}, "gridfxchain/distribution/MsgModifyWithdrawAddress", nil)
	cdc.RegisterConcrete(CommunityPoolSpendProposal{}, "gridfxchain/distribution/CommunityPoolSpendProposal", nil)
	cdc.RegisterConcrete(ChangeDistributionTypeProposal{}, "gridfxchain/distribution/ChangeDistributionTypeProposal", nil)
	cdc.RegisterConcrete(WithdrawRewardEnabledProposal{}, "gridfxchain/distribution/WithdrawRewardEnabledProposal", nil)
	cdc.RegisterConcrete(RewardTruncatePrecisionProposal{}, "gridfxchain/distribution/RewardTruncatePrecisionProposal", nil)
	cdc.RegisterConcrete(MsgWithdrawDelegatorAllRewards{}, "gridfxchain/distribution/MsgWithdrawDelegatorAllRewards", nil)
}

// ModuleCdc generic sealed codec to be used throughout module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
