package gobotapi

import "gobotapi/types"

func (ctx *Client) Self() types.User {
	return *ctx.me
}
