package client

import (
	"github.com/gridfx/fxchain/x/dex/client/cli"
	"github.com/gridfx/fxchain/x/dex/client/rest"
	govclient "github.com/gridfx/fxchain/x/gov/client"
)

// param change proposal handler
var (
	// DelistProposalHandler alias gov NewProposalHandler
	DelistProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitDelistProposal, rest.DelistProposalRESTHandler)
)
