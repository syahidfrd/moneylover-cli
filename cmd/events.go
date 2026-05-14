package cmd

import (
	"github.com/spf13/cobra"
)

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Manage events",
}

var eventsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all events",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAction("list"); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.EventListAll()
		if err != nil {
			return err
		}
		return outputJSON(data)
	},
}

func init() {
	eventsCmd.AddCommand(eventsListCmd)
	rootCmd.AddCommand(eventsCmd)
}
