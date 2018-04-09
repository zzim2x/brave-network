package account

import (
	"github.com/zzim2x/brave-network/cli/command"
	"github.com/spf13/cobra"
	"fmt"
	"github.com/stellar/go/clients/horizon"
)

func NewAccountCommand(cli *command.BraveCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account",
		Short: "account",
		Long:  "",
	}

	balance := &cobra.Command{
		Use:   "balance",
		Short: "balance",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			alias := cmd.Flags().Lookup("alias").Value.String()
			account, err := cli.HorizonClient().LoadAccount(cli.Account[alias])
			if err == nil {
				fmt.Println("My account address:", account.AccountID)
				for _, v := range account.Balances {
					fmt.Printf("type: %s balance: %s\n", v.Asset.Type, v.Balance)
				}
			} else {
				horizonErr := err.(*horizon.Error)
				if horizonErr.Response.StatusCode == 404 {
					fmt.Println("user not found")
				}
				panic(err)
			}
		},
	}

	balance.Flags().String("alias", "", "")
	balance.MarkFlagRequired("alias")

	cmd.AddCommand(balance)

	return cmd
}