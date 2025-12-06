package qbittorrent

import "context"

func (c *Client) SessionGet(ctx context.Context) error {
	return post(ctx, c, Request{Method: "session_get"})
}
