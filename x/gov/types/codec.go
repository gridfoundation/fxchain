package types

import (
	"github.com/gridfx/fxchain/libs/cosmos-sdk/codec"
)

// module codec
var ModuleCdc = codec.New()

// RegisterCodec registers all the necessary types and interfaces for
// governance.
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*Content)(nil), nil)

	cdc.RegisterConcrete(MsgSubmitProposal{}, "fxchain/gov/MsgSubmitProposal", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "fxchain/gov/MsgDeposit", nil)
	cdc.RegisterConcrete(MsgVote{}, "fxchain/gov/MsgVote", nil)

	cdc.RegisterConcrete(TextProposal{}, "fxchain/gov/TextProposal", nil)
	cdc.RegisterConcrete(SoftwareUpgradeProposal{}, "fxchain/gov/SoftwareUpgradeProposal", nil)
}

// RegisterProposalTypeCodec registers an external proposal content type defined
// in another module for the internal ModuleCdc. This allows the MsgSubmitProposal
// to be correctly Amino encoded and decoded.
func RegisterProposalTypeCodec(o interface{}, name string) {
	ModuleCdc.RegisterConcrete(o, name, nil)
}

// TODO determine a good place to seal this codec
func init() {
	RegisterCodec(ModuleCdc)
}
