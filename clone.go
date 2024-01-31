package gobotapi

import (
	"errors"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/methods"
)

func (ctx *BasicClient) Clone(token string, dropUpdates bool, maxConnections int) error {
	if len(ctx.cloningURL) == 0 {
		return errors.New("clone not supported")
	}
	c := Client{
		BasicClient: ctx,
		Token:       token,
	}
	request := &methods.SetWebhook{
		AllowedUpdates:     ctx.AllowedUpdates,
		DropPendingUpdates: dropUpdates,
		MaxConnections:     maxConnections,
		URL:                fmt.Sprintf("%s?token=%s", ctx.cloningURL, token),
	}
	_, err := c.Invoke(request)
	return err
}
