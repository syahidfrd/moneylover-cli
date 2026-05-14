package api

import "encoding/json"

func (c *Client) BudgetList() (json.RawMessage, error) {
	return c.Post("/budget/list/all", map[string]any{})
}

func (c *Client) BudgetAdd(categoryID string, amount float64, walletID string, startDate string, endDate string, isRepeat bool) (json.RawMessage, error) {
	body := map[string]any{
		"categoryId": categoryID,
		"amount":     amount,
		"walletId":   walletID,
		"startDate":  startDate,
		"endDate":    endDate,
		"isRepeat":   isRepeat,
	}
	return c.Post("/budget/add", body)
}

func (c *Client) BudgetEdit(budgetID string, categoryID string, amount float64, walletID string, startDate string, endDate string, isRepeat bool) (json.RawMessage, error) {
	body := map[string]any{
		"budgetId":   budgetID,
		"categoryId": categoryID,
		"amount":     amount,
		"walletId":   walletID,
		"startDate":  startDate,
		"endDate":    endDate,
		"isRepeat":   isRepeat,
	}
	return c.Post("/budget/edit", body)
}

func (c *Client) BudgetDelete(budgetID string) (json.RawMessage, error) {
	return c.Post("/budget/delete", map[string]any{"_id": budgetID})
}
