package cmd

import (
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User commands",
}

var userInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get user info",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAction("info"); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.UserInfo()
		if err != nil {
			return err
		}
		return outputJSON(data)
	},
}

func init() {
	userCmd.AddCommand(userInfoCmd)
	rootCmd.AddCommand(userCmd)
}
