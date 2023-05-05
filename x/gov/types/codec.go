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

	cdc.RegisterConcrete(MsgSubmitProposal{}, "gridchain/gov/MsgSubmitProposal", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "gridchain/gov/MsgDeposit", nil)
	cdc.RegisterConcrete(MsgVote{}, "gridchain/gov/MsgVote", nil)

	cdc.RegisterConcrete(TextProposal{}, "gridchain/gov/TextProposal", nil)
	cdc.RegisterConcrete(SoftwareUpgradeProposal{}, "gridchain/gov/SoftwareUpgradeProposal", nil)
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
