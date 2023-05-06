package client

import (
	"github.com/gridfx/fxchain/x/feesplit/client/cli"
	"github.com/gridfx/fxchain/x/feesplit/client/rest"
	govcli "github.com/gridfx/fxchain/x/gov/client"
)

var (
	// FeeSplitSharesProposalHandler alias gov NewProposalHandler
	FeeSplitSharesProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdFeeSplitSharesProposal,
		rest.FeeSplitSharesProposalRESTHandler,
	)
)
