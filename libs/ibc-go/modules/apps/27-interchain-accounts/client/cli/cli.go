package cli

import (
	"github.com/gridfx/fxchain/libs/cosmos-sdk/codec"
	interfacetypes "github.com/gridfx/fxchain/libs/cosmos-sdk/codec/types"
	controllercli "github.com/gridfx/fxchain/libs/ibc-go/modules/apps/27-interchain-accounts/controller/client/cli"
	hostcli "github.com/gridfx/fxchain/libs/ibc-go/modules/apps/27-interchain-accounts/host/client/cli"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the query commands for the interchain-accounts submodule
func GetQueryCmd(cdc *codec.CodecProxy, reg interfacetypes.InterfaceRegistry) *cobra.Command {
	icaQueryCmd := &cobra.Command{
		Use:                        "interchain-accounts",
		Aliases:                    []string{"ica"},
		Short:                      "interchain-accounts subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	icaQueryCmd.AddCommand(
		controllercli.GetQueryCmd(cdc, reg),
		hostcli.GetQueryCmd(cdc, reg),
	)

	return icaQueryCmd
}
