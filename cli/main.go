package main

import (
	"github.com/spf13/cobra"
	"github.com/jinzhu/configor"
	"github.com/zzim2x/brave-network/cli/command"
	"github.com/zzim2x/brave-network/cli/command/commands"
)

func main() {
	braveCli := &command.BraveCli{}
	configor.Load(&braveCli, "config.yaml")

	cmd := newCliCommand(braveCli)
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}

func newCliCommand(cli *command.BraveCli) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "brave",
		Short: "brave",
		Long:  "",
	}

	commands.AddCommands(rootCmd, cli)
	return rootCmd
}
