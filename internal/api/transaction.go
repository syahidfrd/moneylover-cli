package api

import "encoding/json"

func (c *Client) TransactionList(walletID string, startDate string, endDate string) (json.RawMessage, error) {
	body := map[string]any{
		"walletId":  walletID,
		"startDate": startDate,
		"endDate":   endDate,
	}
	return c.Post("/transaction/list", body)
}

func (c *Client) TransactionAdd(wallet string, category string, amount float64, note string, date string) (json.RawMessage, error) {
	body := map[string]any{
		"with":           []string{},
		"account":        wallet,
		"category":       category,
		"amount":         amount,
		"note":           note,
		"displayDate":    date,
		"event":          "",
		"exclude_report": false,
		"longtitude":     0,
		"latitude":       0,
		"image":          "",
	}
	return c.Post("/transaction/add", body)
}

func (c *Client) TransactionDebts(wallets []string) (json.RawMessage, error) {
	return c.Post("/transaction/debts", map[string]any{
		"accounts": wallets,
	})
}
