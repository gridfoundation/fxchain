package client

import (
	"github.com/gridfx/fxchain/libs/cosmos-sdk/x/mint/client/cli"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/x/mint/client/rest"
	govcli "github.com/gridfx/fxchain/x/gov/client"
)

var (
	ManageTreasuresProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdManageTreasuresProposal,
		rest.ManageTreasuresProposalRESTHandler,
	)
	ModifyNextBlockUpdateProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdModifyNextBlockUpdateProposal,
		rest.ModifyNextBlockUpdateProposalRESTHandler,
	)
)
