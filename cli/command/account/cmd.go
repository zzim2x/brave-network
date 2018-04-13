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
			address := cmd.Flags().Lookup("address").Value.String()

			if alias != "" {
				address = cli.Account[alias]
			} else if address == "" {
				cmd.Usage()
				return
			}

			account, err := cli.HorizonClient().LoadAccount(address)

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

	balance.Flags().String("alias", "", "config.yaml 설정된 계정 alias")
	balance.Flags().String("address", "", "계좌 주소")

	cmd.AddCommand(balance)

	return cmd
}