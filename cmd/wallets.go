package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/syahidfrd/moneylover-cli/internal/api"
	"github.com/syahidfrd/moneylover-cli/internal/config"
)

var walletsCmd = &cobra.Command{
	Use:   "wallets",
	Short: "Manage wallets",
}

var walletsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all wallets",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAction("list"); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.WalletList()
		if err != nil {
			return err
		}
		return outputJSON(data)
	},
}

var walletsAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new wallet",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAction("add"); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		name, _ := cmd.Flags().GetString("name")
		currencyID, _ := cmd.Flags().GetInt("currency-id")
		icon, _ := cmd.Flags().GetString("icon")
		accountType, _ := cmd.Flags().GetInt("account-type")
		balance, _ := cmd.Flags().GetFloat64("balance")

		hasBalance := cmd.Flags().Changed("balance")
		data, err := client.WalletAdd(name, currencyID, icon, accountType, hasBalance, balance)
		if err != nil {
			return err
		}
		return outputJSON(data)
	},
}

var walletsEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a wallet",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAction("edit"); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		id, _ := cmd.Flags().GetString("id")
		name, _ := cmd.Flags().GetString("name")
		icon, _ := cmd.Flags().GetString("icon")
		currencyID, _ := cmd.Flags().GetInt("currency-id")
		accountType, _ := cmd.Flags().GetInt("account-type")

		data, err := client.WalletEdit(id, name, icon, currencyID, accountType)
		if err != nil {
			return err
		}
		return outputJSON(data)
	},
}

var walletsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a wallet",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAction("delete"); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		id, _ := cmd.Flags().GetString("id")
		data, err := client.WalletDelete(id)
		if err != nil {
			return err
		}
		return outputJSON(data)
	},
}

func newClient() (*api.Client, error) {
	path := getTokenPath()
	token, err := config.Load(path)
	if err != nil {
		return nil, fmt.Errorf("not authenticated. Run 'moneylover auth login' first")
	}
	return api.NewClient(token), nil
}

func outputJSON(data json.RawMessage) error {
	out := map[string]json.RawMessage{"data": data}
	output, _ := json.MarshalIndent(out, "", "  ")
	fmt.Fprintln(os.Stdout, string(output))
	return nil
}

func init() {
	walletsAddCmd.Flags().String("name", "", "wallet name (required)")
	walletsAddCmd.Flags().Int("currency-id", 44, "currency ID (default: 44/IDR)")
	walletsAddCmd.Flags().String("icon", "icon_80", "wallet icon")
	walletsAddCmd.Flags().Int("account-type", 0, "account type (0=normal, 2=linked, 4=credit card, 5=saving)")
	walletsAddCmd.Flags().Float64("balance", 0, "initial balance")
	walletsAddCmd.MarkFlagRequired("name")

	walletsEditCmd.Flags().String("id", "", "wallet ID (required)")
	walletsEditCmd.Flags().String("name", "", "wallet name (required)")
	walletsEditCmd.Flags().String("icon", "icon_80", "wallet icon")
	walletsEditCmd.Flags().Int("currency-id", 44, "currency ID")
	walletsEditCmd.Flags().Int("account-type", 0, "account type")
	walletsEditCmd.MarkFlagRequired("id")
	walletsEditCmd.MarkFlagRequired("name")

	walletsDeleteCmd.Flags().String("id", "", "wallet ID (required)")
	walletsDeleteCmd.MarkFlagRequired("id")

	walletsCmd.AddCommand(walletsListCmd)
	walletsCmd.AddCommand(walletsAddCmd)
	walletsCmd.AddCommand(walletsEditCmd)
	walletsCmd.AddCommand(walletsDeleteCmd)
	rootCmd.AddCommand(walletsCmd)
}
