package rest

import (
	"github.com/gorilla/mux"
	govRest "github.com/gridfx/fxchain/x/gov/client/rest"

	"github.com/gridfx/fxchain/libs/cosmos-sdk/client/context"
)

// RegisterRoutes registers farm-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}

// ManageWhiteListProposalRESTHandler defines farm proposal handler
func ManageWhiteListProposalRESTHandler(context.CLIContext) govRest.ProposalRESTHandler {
	return govRest.ProposalRESTHandler{}
}
