package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/syahidfrd/moneylover-cli/internal/config"
)

var tokenPath string
var configPath string

var rootCmd = &cobra.Command{
	Use:   "moneylover",
	Short: "Money Lover CLI - unofficial API client",
	Long:  "A command-line interface for the Money Lover personal finance app API.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, `{"error":"%s"}`+"\n", err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&tokenPath, "token-path", "", "path to token file (default: ~/.config/moneylover/token.json)")
	rootCmd.PersistentFlags().StringVar(&configPath, "config-path", "", "path to config file (default: ~/.config/moneylover/config.json)")
}

func getTokenPath() string {
	if tokenPath != "" {
		return tokenPath
	}
	if env := os.Getenv("MONEYLOVER_TOKEN_PATH"); env != "" {
		return env
	}
	home, _ := os.UserHomeDir()
	return home + "/.config/moneylover/token.json"
}

func getConfigPath() string {
	if configPath != "" {
		return configPath
	}
	if env := os.Getenv("MONEYLOVER_CONFIG_PATH"); env != "" {
		return env
	}
	home, _ := os.UserHomeDir()
	return home + "/.config/moneylover/config.json"
}

func checkAction(action string) error {
	cfg, err := config.LoadAppConfig(getConfigPath())
	if err != nil {
		return err
	}
	if !cfg.IsActionAllowed(action) {
		return fmt.Errorf("action '%s' is not allowed by config", action)
	}
	return nil
}