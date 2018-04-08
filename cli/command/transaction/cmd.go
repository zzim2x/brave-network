package transaction

import (
	"github.com/zzim2x/brave-network/cli/command"
	"github.com/spf13/cobra"
	"fmt"
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
	"net/http"
)

func NewTransactionCommand(cli *command.BraveCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transaction",
		Short: "transaction",
		Long:  "",
	}

	client := &horizon.Client{
		URL: cli.Network.Horizon,
		HTTP: http.DefaultClient,
	}

	createAccount := &cobra.Command{
		Use:   "create",
		Short: "create",
		Long:  "create account with initial balance",
		Run: func(cmd *cobra.Command, args []string) {
			seed := cmd.Flags().Lookup("seed").Value.String()
			address := cmd.Flags().Lookup("address").Value.String()
			amount := cmd.Flags().Lookup("amount").Value.String()

			tx, err := build.Transaction(
				build.Network{cli.Network.Passphrase},
				build.SourceAccount{seed},
				build.AutoSequence{client},
				build.CreateAccount(
					build.Destination{address},
					build.NativeAmount{amount},
				),
			)

			if err != nil {
				panic(err)
			}

			submit(client, tx, seed)
		},
	}

	createAccount.Flags().String("seed", "", "")
	createAccount.Flags().String("address", "", "")
	createAccount.Flags().String("amount", "", "")
	createAccount.MarkFlagRequired("seed")
	createAccount.MarkFlagRequired("address")
	createAccount.MarkFlagRequired("amount")

	payment := &cobra.Command{
		Use:   "payment",
		Short: "payment",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			seed := cmd.Flags().Lookup("seed").Value.String()
			address := cmd.Flags().Lookup("address").Value.String()
			amount := cmd.Flags().Lookup("amount").Value.String()

			tx, err := build.Transaction(
				build.Network{cli.Network.Passphrase},
				build.SourceAccount{seed},
				build.AutoSequence{client},
				build.Payment(
					build.Destination{address},
					build.NativeAmount{amount},
				),
			)

			if err != nil {
				panic(err)
			}

			submit(client, tx, seed)
		},
	}

	payment.Flags().String("seed", "", "")
	payment.Flags().String("address", "", "")
	payment.Flags().String("amount", "", "")
	payment.MarkFlagRequired("seed")
	payment.MarkFlagRequired("address")
	payment.MarkFlagRequired("amount")

	cmd.AddCommand(createAccount, payment)

	return cmd
}

func submit(client *horizon.Client, tx *build.TransactionBuilder, seed string) {
	if txe, err := tx.Sign(seed); err == nil {
		if blob, err := txe.Base64(); err == nil {
			if resp, err := client.SubmitTransaction(blob); err == nil {
				fmt.Println("transaction posted in ledger:", resp.Ledger)
			} else {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
}