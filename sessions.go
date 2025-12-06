package transmission

import "context"

func (c *Client) SessionGet(ctx context.Context) error {
	return c.post(ctx, Request{Method: "session_get"})
}
