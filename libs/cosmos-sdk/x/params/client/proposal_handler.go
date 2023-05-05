package client

import (
	govclient "github.com/gridfx/fxchain/libs/cosmos-sdk/x/gov/client"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/x/params/client/cli"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/x/params/client/rest"
)

// ProposalHandler handles param change proposals
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
