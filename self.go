package gobotapi

import "github.com/Squirrel-Network/gobotapi/types"

func (ctx *Client) Self() types.User {
	return *ctx.me
}
