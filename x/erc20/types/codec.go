package types

import (
	"github.com/gridfx/fxchain/libs/cosmos-sdk/codec"
)

// ModuleCdc defines the erc20 module's codec
var ModuleCdc = codec.New()

const (
	TokenMappingProposalName          = "gridfxchain/erc20/TokenMappingProposal"
	ProxyContractRedirectProposalName = "gridfxchain/erc20/ProxyContractRedirectProposal"
	ContractTemplateProposalName      = "gridfxchain/erc20/ContractTemplateProposal"
	CompiledContractProposalName      = "gridfxchain/erc20/Contract"
)

// RegisterCodec registers all the necessary types and interfaces for the
// erc20 module
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(TokenMappingProposal{}, TokenMappingProposalName, nil)

	cdc.RegisterConcrete(ProxyContractRedirectProposal{}, ProxyContractRedirectProposalName, nil)
	cdc.RegisterConcrete(ContractTemplateProposal{}, ContractTemplateProposalName, nil)
	cdc.RegisterConcrete(CompiledContract{}, CompiledContractProposalName, nil)
}

func init() {
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
