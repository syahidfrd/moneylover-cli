package api

import "encoding/json"

func (c *Client) CategoryListAll() (json.RawMessage, error) {
	return c.Post("/category/list-all", map[string]any{})
}

func (c *Client) CategoryAdd(name string, categoryType int, icon string, wallet string) (json.RawMessage, error) {
	body := map[string]any{
		"name":    name,
		"type":    categoryType,
		"icon":    icon,
		"account": wallet,
	}
	return c.Post("/category/add", body)
}

func (c *Client) CategoryEdit(id string, name string, icon string) (json.RawMessage, error) {
	body := map[string]any{
		"_id":  id,
		"name": name,
		"icon": icon,
	}
	return c.Post("/category/edit", body)
}

func (c *Client) CategoryDelete(id string) (json.RawMessage, error) {
	return c.Post("/category/delete", map[string]any{"_id": id})
}
