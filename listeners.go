package gobotapi

import "github.com/GoBotApiOfficial/gobotapi/types"

func (ctx *BasicClient) OnAnyMessageEvent(handler func(client *Client, update types.Message)) {
	ctx.OnMessage(handler)
	ctx.OnEditedMessage(handler)
	ctx.OnChannelPost(handler)
	ctx.OnEditedChannelPost(handler)
}

func (ctx *BasicClient) OnAnyMessage(handler func(client *Client, update types.Message)) {
	ctx.OnMessage(handler)
	ctx.OnChannelPost(handler)
}

func (ctx *BasicClient) OnAnyEditedMessage(handler func(client *Client, update types.Message)) {
	ctx.OnEditedMessage(handler)
	ctx.OnEditedChannelPost(handler)
}
