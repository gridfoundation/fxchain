package types

import (
	"github.com/gridfx/fxchain/libs/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgWithdrawValidatorCommission{}, "okfxchain/distribution/MsgWithdrawReward", nil)
	cdc.RegisterConcrete(MsgWithdrawDelegatorReward{}, "okfxchain/distribution/MsgWithdrawDelegatorReward", nil)
	cdc.RegisterConcrete(MsgSetWithdrawAddress{}, "okfxchain/distribution/MsgModifyWithdrawAddress", nil)
	cdc.RegisterConcrete(CommunityPoolSpendProposal{}, "okfxchain/distribution/CommunityPoolSpendProposal", nil)
	cdc.RegisterConcrete(ChangeDistributionTypeProposal{}, "okfxchain/distribution/ChangeDistributionTypeProposal", nil)
	cdc.RegisterConcrete(WithdrawRewardEnabledProposal{}, "okfxchain/distribution/WithdrawRewardEnabledProposal", nil)
	cdc.RegisterConcrete(RewardTruncatePrecisionProposal{}, "okfxchain/distribution/RewardTruncatePrecisionProposal", nil)
	cdc.RegisterConcrete(MsgWithdrawDelegatorAllRewards{}, "okfxchain/distribution/MsgWithdrawDelegatorAllRewards", nil)
}

// ModuleCdc generic sealed codec to be used throughout module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
