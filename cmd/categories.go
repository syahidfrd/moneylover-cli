package cmd

import (
	"github.com/spf13/cobra"
)

var categoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Manage categories",
}

var categoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all categories",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAction("list"); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.CategoryListAll()
		if err != nil {
			return err
		}
		return outputJSON(data)
	},
}

var categoriesAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new category",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAction("add"); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		name, _ := cmd.Flags().GetString("name")
		categoryType, _ := cmd.Flags().GetInt("type")
		icon, _ := cmd.Flags().GetString("icon")
		wallet, _ := cmd.Flags().GetString("wallet")

		data, err := client.CategoryAdd(name, categoryType, icon, wallet)
		if err != nil {
			return err
		}
		return outputJSON(data)
	},
}

var categoriesEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a category",
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

		data, err := client.CategoryEdit(id, name, icon)
		if err != nil {
			return err
		}
		return outputJSON(data)
	},
}

var categoriesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a category",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAction("delete"); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		id, _ := cmd.Flags().GetString("id")
		data, err := client.CategoryDelete(id)
		if err != nil {
			return err
		}
		return outputJSON(data)
	},
}

func init() {
	categoriesAddCmd.Flags().String("name", "", "category name (required)")
	categoriesAddCmd.Flags().Int("type", 2, "category type (1=income, 2=expense)")
	categoriesAddCmd.Flags().String("icon", "ic_category_other", "category icon")
	categoriesAddCmd.Flags().String("wallet", "", "wallet ID (required)")
	categoriesAddCmd.MarkFlagRequired("name")
	categoriesAddCmd.MarkFlagRequired("wallet")

	categoriesEditCmd.Flags().String("id", "", "category ID (required)")
	categoriesEditCmd.Flags().String("name", "", "category name (required)")
	categoriesEditCmd.Flags().String("icon", "", "category icon")
	categoriesEditCmd.MarkFlagRequired("id")
	categoriesEditCmd.MarkFlagRequired("name")

	categoriesDeleteCmd.Flags().String("id", "", "category ID (required)")
	categoriesDeleteCmd.MarkFlagRequired("id")

	categoriesCmd.AddCommand(categoriesListCmd)
	categoriesCmd.AddCommand(categoriesAddCmd)
	categoriesCmd.AddCommand(categoriesEditCmd)
	categoriesCmd.AddCommand(categoriesDeleteCmd)
	rootCmd.AddCommand(categoriesCmd)
}
