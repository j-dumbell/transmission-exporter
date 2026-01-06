package transmission

import "context"

func (c *Client) QueueMoveTop(ctx context.Context, ids *TorrentIDs) error {
	_, err := postWithArgs[torrentIDsArgs, any](ctx, c, "queue-move-top", torrentIDsArgs{ids})
	return err
}

func (c *Client) QueueMoveUp(ctx context.Context, ids *TorrentIDs) error {
	_, err := postWithArgs[torrentIDsArgs, any](ctx, c, "queue-move-up", torrentIDsArgs{ids})
	return err
}

func (c *Client) QueueMoveDown(ctx context.Context, ids *TorrentIDs) error {
	_, err := postWithArgs[torrentIDsArgs, any](ctx, c, "queue-move-down", torrentIDsArgs{ids})
	return err
}

func (c *Client) QueueMoveBottom(ctx context.Context, ids *TorrentIDs) error {
	_, err := postWithArgs[torrentIDsArgs, any](ctx, c, "queue-move-bottom", torrentIDsArgs{ids})
	return err
}
