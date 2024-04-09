package gobotapi

import "github.com/GoBotApiOfficial/gobotapi/types"

func (ctx *Client) Self() types.User {
	return *ctx.me
}
