package api

import "encoding/json"

func (c *Client) WalletList() (json.RawMessage, error) {
	return c.Post("/wallet/list", map[string]any{})
}

func (c *Client) WalletAdd(name string, currencyID int, icon string, accountType int, hasBalance bool, amount float64) (json.RawMessage, error) {
	body := map[string]any{
		"account_type":  accountType,
		"name":          name,
		"exclude_total": false,
		"icon":          icon,
		"currency_id":   currencyID,
	}
	if hasBalance {
		body["hasBalance"] = true
		body["transaction"] = map[string]any{
			"amount":       amount,
			"typeCategory": 1,
			"note":         "Initial balance",
		}
	}
	return c.Post("/wallet/add", body)
}
