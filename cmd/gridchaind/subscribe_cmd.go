package main

import (
	"fmt"
	"github.com/gridfx/fxchain/app/logevents"
	"github.com/gridfx/fxchain/libs/system"
	"github.com/spf13/cobra"

	"github.com/gridfx/fxchain/libs/cosmos-sdk/codec"
)

func subscribeCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscribe",
		Short: "subscribe " + system.ChainName + " logs from kafka",
	}
	cmd.AddCommand(subscribeLog())
	return cmd
}

func subscribeLog() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logs [urls] [outdir]",
		Short: "logs [urls] [outdir]",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s, %s\n", args[0], args[1])
			subscriber := logevents.NewSubscriber()
			subscriber.Init(args[0], args[1])
			subscriber.Run()
		},
	}
	return cmd
}
