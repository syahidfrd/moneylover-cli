package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/syahidfrd/moneylover-cli/internal/api"
	"github.com/syahidfrd/moneylover-cli/internal/config"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication commands",
}

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Get OAuth login URL",
	RunE: func(cmd *cobra.Command, args []string) error {
		loginURL, err := api.GetLoginURL()
		if err != nil {
			return err
		}

		output, _ := json.Marshal(map[string]string{"login_url": loginURL})
		fmt.Println(string(output))
		return nil
	},
}

var authCallbackCmd = &cobra.Command{
	Use:   "callback [redirect_url]",
	Short: "Save token from OAuth callback URL",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		token, err := parseCallbackURL(args[0])
		if err != nil {
			return err
		}

		path := getTokenPath()
		if err := config.Save(path, token); err != nil {
			return err
		}

		fmt.Fprintf(os.Stderr, "Token saved to %s\n", path)
		output, _ := json.Marshal(map[string]string{"status": "login_success"})
		fmt.Println(string(output))
		return nil
	},
}

var authStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show authentication status",
	RunE: func(cmd *cobra.Command, args []string) error {
		path := getTokenPath()
		token, err := config.Load(path)
		if err != nil {
			return fmt.Errorf("not authenticated: %w", err)
		}

		claims, err := decodeJWTClaims(token.AccessToken)
		if err != nil {
			return fmt.Errorf("failed to decode token: %w", err)
		}

		expTime := time.Unix(token.Expire, 0)
		isValid := time.Now().Before(expTime)

		status := map[string]any{
			"user_id":    claims["userId"],
			"client":     claims["client"],
			"expires_at": expTime.Format(time.RFC3339),
			"valid":      isValid,
		}

		output, _ := json.MarshalIndent(status, "", "  ")
		fmt.Println(string(output))
		return nil
	},
}

func parseCallbackURL(rawURL string) (*config.Token, error) {
	// URL format: https://web.moneylover.me/login?access_token=...&refresh_token=...&expire=...&status=true
	parts := strings.SplitN(rawURL, "?", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid redirect URL")
	}

	params := make(map[string]string)
	for _, pair := range strings.Split(parts[1], "&") {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) == 2 {
			params[kv[0]] = kv[1]
		}
	}

	accessToken := params["access_token"]
	refreshToken := params["refresh_token"]
	expireStr := params["expire"]

	if accessToken == "" || refreshToken == "" {
		return nil, fmt.Errorf("missing access_token or refresh_token in URL")
	}

	var expire int64
	fmt.Sscanf(expireStr, "%d", &expire)

	return &config.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expire:       expire,
	}, nil
}

func decodeJWTClaims(token string) (map[string]any, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid JWT format")
	}

	payload := parts[1]
	// Add padding if needed
	switch len(payload) % 4 {
	case 2:
		payload += "=="
	case 3:
		payload += "="
	}

	decoded, err := base64.URLEncoding.DecodeString(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JWT payload: %w", err)
	}

	var claims map[string]any
	if err := json.Unmarshal(decoded, &claims); err != nil {
		return nil, fmt.Errorf("failed to parse JWT claims: %w", err)
	}

	return claims, nil
}

func init() {
	authCmd.AddCommand(authLoginCmd)
	authCmd.AddCommand(authCallbackCmd)
	authCmd.AddCommand(authStatusCmd)
	rootCmd.AddCommand(authCmd)
}
