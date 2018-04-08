package keypair

import (
	"github.com/zzim2x/brave-network/cli/command"
	"github.com/spf13/cobra"
	"fmt"
	"github.com/stellar/go/keypair"
)

func NewKeyPairCommand(cli *command.BraveCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keypair",
		Short: "keypair",
		Long:  "",
	}

	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "generate seed",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			pair, err := keypair.Random()
			if err != nil {
				panic(err)
			}
			fmt.Println("Seed:", pair.Seed())
		},
	}

	parseCmd := &cobra.Command{
		Use:   "parse",
		Short: "parse seed",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			seed := cmd.Flags().Lookup("seed").Value.String()
			pair, err := keypair.Parse(seed)

			if err != nil {
				panic(err)
			}

			fmt.Println("Address:", pair.Address())
		},
	}

	parseCmd.Flags().String("seed", "", "")
	parseCmd.MarkFlagRequired("seed")

	cmd.AddCommand(
		generateCmd,
		parseCmd,
	)

	return cmd
}
