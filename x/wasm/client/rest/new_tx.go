package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	clientCtx "github.com/gridfx/fxchain/libs/cosmos-sdk/client/context"
	"github.com/gridfx/fxchain/libs/cosmos-sdk/types/rest"
	"github.com/gridfx/fxchain/x/wasm/client/utils"
	"github.com/gridfx/fxchain/x/wasm/types"
)

func registerNewTxRoutes(cliCtx clientCtx.CLIContext, r *mux.Router) {
	r.HandleFunc("/wasm/contract/{contractAddr}/admin", setContractAdminHandlerFn(cliCtx)).Methods("PUT")
	r.HandleFunc("/wasm/contract/{contractAddr}/code", migrateContractHandlerFn(cliCtx)).Methods("PUT")
}

type migrateContractReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
	Admin   string       `json:"admin,omitempty" yaml:"admin"`
	CodeID  uint64       `json:"code_id" yaml:"code_id"`
	Msg     []byte       `json:"msg,omitempty" yaml:"msg"`
}

type updateContractAdministrateReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
	Admin   string       `json:"admin,omitempty" yaml:"admin"`
}

func setContractAdminHandlerFn(cliCtx clientCtx.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req updateContractAdministrateReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}
		vars := mux.Vars(r)
		contractAddr := vars["contractAddr"]

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := &types.MsgUpdateAdmin{
			Sender:   req.BaseReq.From,
			NewAdmin: req.Admin,
			Contract: contractAddr,
		}
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGeneratedTxResponse(cliCtx, w, req.BaseReq, msg)
	}
}

func migrateContractHandlerFn(cliCtx clientCtx.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req migrateContractReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}
		vars := mux.Vars(r)
		contractAddr := vars["contractAddr"]

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := &types.MsgMigrateContract{
			Sender:   req.BaseReq.From,
			Contract: contractAddr,
			CodeID:   req.CodeID,
			Msg:      req.Msg,
		}
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGeneratedTxResponse(cliCtx, w, req.BaseReq, msg)
	}
}
