package client

import (
	govclient "github.com/gridfx/fxchain/libs/cosmos-sdk/x/gov/client"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/x/upgrade/client/cli"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
