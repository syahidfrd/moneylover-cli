package cmd

import (
	"github.com/spf13/cobra"
)

var budgetsCmd = &cobra.Command{
	Use:   "budgets",
	Short: "Manage budgets",
}

var budgetsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all budgets",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.BudgetList()
		if err != nil {
			return err
		}
		return outputJSON(data)
	},
}

var budgetsAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new budget",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		category, _ := cmd.Flags().GetString("category")
		amount, _ := cmd.Flags().GetFloat64("amount")
		wallet, _ := cmd.Flags().GetString("wallet")
		start, _ := cmd.Flags().GetString("start")
		end, _ := cmd.Flags().GetString("end")
		repeat, _ := cmd.Flags().GetBool("repeat")

		data, err := client.BudgetAdd(category, amount, wallet, start, end, repeat)
		if err != nil {
			return err
		}
		return outputJSON(data)
	},
}

func init() {
	budgetsAddCmd.Flags().String("category", "", "category ID (required)")
	budgetsAddCmd.Flags().Float64("amount", 0, "budget amount (required)")
	budgetsAddCmd.Flags().String("wallet", "", "wallet ID (required)")
	budgetsAddCmd.Flags().String("start", "", "start date YYYY-MM-DD (required)")
	budgetsAddCmd.Flags().String("end", "", "end date YYYY-MM-DD (required)")
	budgetsAddCmd.Flags().Bool("repeat", false, "repeat monthly")
	budgetsAddCmd.MarkFlagRequired("category")
	budgetsAddCmd.MarkFlagRequired("amount")
	budgetsAddCmd.MarkFlagRequired("wallet")
	budgetsAddCmd.MarkFlagRequired("start")
	budgetsAddCmd.MarkFlagRequired("end")

	budgetsCmd.AddCommand(budgetsListCmd)
	budgetsCmd.AddCommand(budgetsAddCmd)
	rootCmd.AddCommand(budgetsCmd)
}
