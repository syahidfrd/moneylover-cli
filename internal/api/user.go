package api

import "encoding/json"

func (c *Client) EventListAll() (json.RawMessage, error) {
	return c.Post("/event/list/all", map[string]any{})
}

func (c *Client) UserInfo() (json.RawMessage, error) {
	return c.Post("/user/info", map[string]any{})
}
