package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
	"bytes"
)

func GetLoginURL() (string, error) {
	body := []byte(`{}`)
	req, err := http.NewRequest("POST", BaseURL+"/user/login-url", bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	var apiResp struct {
		Error int    `json:"error"`
		Msg   string `json:"msg"`
		Data  struct {
			LoginURL string `json:"login_url"`
		} `json:"data"`
	}
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if apiResp.Error != 0 {
		return "", fmt.Errorf("API error: %s", apiResp.Msg)
	}

	return apiResp.Data.LoginURL, nil
}

func (c *Client) RefreshToken() error {
	body := map[string]string{"refresh_token": c.token.RefreshToken}
	_, err := c.PostRaw("/user/refresh-token", body)
	return err
}
