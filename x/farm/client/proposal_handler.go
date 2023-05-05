package client

import (
	"github.com/gridfx/fxchain/x/farm/client/cli"
	"github.com/gridfx/fxchain/x/farm/client/rest"
	govcli "github.com/gridfx/fxchain/x/gov/client"
)

var (
	// ManageWhiteListProposalHandler alias gov NewProposalHandler
	ManageWhiteListProposalHandler = govcli.NewProposalHandler(cli.GetCmdManageWhiteListProposal, rest.ManageWhiteListProposalRESTHandler)
)
