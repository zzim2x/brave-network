package account

import (
	"github.com/zzim2x/brave-network/cli/command"
	"github.com/spf13/cobra"
)

func NewAccountCommand(cli *command.BraveCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account",
		Short: "account",
		Long:  "",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "balance",
		Short: "balance",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
		},
	})

	return cmd
}