package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/syahidfrd/moneylover-cli/internal/config"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage CLI configuration",
}

var configSetAllowedActionsCmd = &cobra.Command{
	Use:   "set-allowed-actions [actions]",
	Short: "Set allowed actions (comma-separated: list,add,edit,delete,debts,info,login,callback,status)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		actions := strings.Split(args[0], ",")
		for i := range actions {
			actions[i] = strings.TrimSpace(actions[i])
		}

		path := getConfigPath()
		cfg, err := config.LoadAppConfig(path)
		if err != nil {
			return err
		}

		cfg.AllowedActions = actions
		if err := config.SaveAppConfig(path, cfg); err != nil {
			return err
		}

		output, _ := json.Marshal(map[string]any{
			"allowed_actions": actions,
		})
		fmt.Println(string(output))
		return nil
	},
}

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Show current configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		path := getConfigPath()
		cfg, err := config.LoadAppConfig(path)
		if err != nil {
			return err
		}

		output, _ := json.MarshalIndent(cfg, "", "  ")
		fmt.Println(string(output))
		return nil
	},
}

var configResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset configuration (allow all actions)",
	RunE: func(cmd *cobra.Command, args []string) error {
		path := getConfigPath()
		cfg := &config.AppConfig{}
		if err := config.SaveAppConfig(path, cfg); err != nil {
			return err
		}

		output, _ := json.Marshal(map[string]string{"status": "config_reset"})
		fmt.Println(string(output))
		return nil
	},
}

func init() {
	configCmd.AddCommand(configSetAllowedActionsCmd)
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(configResetCmd)
	rootCmd.AddCommand(configCmd)
}
