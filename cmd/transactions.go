package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var transactionsCmd = &cobra.Command{
	Use:   "transactions",
	Short: "Manage transactions",
}

var transactionsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List transactions",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		wallet, _ := cmd.Flags().GetString("wallet")
		start, _ := cmd.Flags().GetString("start")
		end, _ := cmd.Flags().GetString("end")

		data, err := client.TransactionList(wallet, start, end)
		if err != nil {
			return err
		}
		return outputJSON(data)
	},
}

var transactionsAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new transaction",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		wallet, _ := cmd.Flags().GetString("wallet")
		category, _ := cmd.Flags().GetString("category")
		amount, _ := cmd.Flags().GetFloat64("amount")
		note, _ := cmd.Flags().GetString("note")
		date, _ := cmd.Flags().GetString("date")

		data, err := client.TransactionAdd(wallet, category, amount, note, date)
		if err != nil {
			return err
		}
		return outputJSON(data)
	},
}

var transactionsDebtsCmd = &cobra.Command{
	Use:   "debts",
	Short: "List debts/loans",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		walletsStr, _ := cmd.Flags().GetString("wallets")
		wallets := strings.Split(walletsStr, ",")

		data, err := client.TransactionDebts(wallets)
		if err != nil {
			return err
		}
		return outputJSON(data)
	},
}

func init() {
	transactionsListCmd.Flags().String("wallet", "all", "wallet ID or 'all'")
	transactionsListCmd.Flags().String("start", "", "start date (YYYY-MM-DD) (required)")
	transactionsListCmd.Flags().String("end", "", "end date (YYYY-MM-DD) (required)")
	transactionsListCmd.MarkFlagRequired("start")
	transactionsListCmd.MarkFlagRequired("end")

	transactionsAddCmd.Flags().String("wallet", "", "wallet ID (required)")
	transactionsAddCmd.Flags().String("category", "", "category ID (required)")
	transactionsAddCmd.Flags().Float64("amount", 0, "amount (required)")
	transactionsAddCmd.Flags().String("note", "", "transaction note")
	transactionsAddCmd.Flags().String("date", "", "display date YYYY-MM-DD (required)")
	transactionsAddCmd.MarkFlagRequired("wallet")
	transactionsAddCmd.MarkFlagRequired("category")
	transactionsAddCmd.MarkFlagRequired("amount")
	transactionsAddCmd.MarkFlagRequired("date")

	transactionsDebtsCmd.Flags().String("wallets", "", "comma-separated wallet IDs (required)")
	transactionsDebtsCmd.MarkFlagRequired("wallets")

	transactionsCmd.AddCommand(transactionsListCmd)
	transactionsCmd.AddCommand(transactionsAddCmd)
	transactionsCmd.AddCommand(transactionsDebtsCmd)
	rootCmd.AddCommand(transactionsCmd)
}
