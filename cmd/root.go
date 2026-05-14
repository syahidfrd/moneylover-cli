package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var tokenPath string

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