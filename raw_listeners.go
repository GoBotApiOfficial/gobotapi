package gobotapi

import "github.com/Squirrel-Network/gobotapi/types"


func (ctx *Client) OnRawUpdate(handler func(update types.Update)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]interface{})
	}
	ctx.handlers["raw"] = append(ctx.handlers["raw"], handler)
}

func (ctx *Client) OnMessage(handler func(update types.Message)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]interface{})
	}
	ctx.handlers["message"] = append(ctx.handlers["message"], handler)
}

func (ctx *Client) OnEditedMessage(handler func(update types.Message)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]interface{})
	}
	ctx.handlers["edited_message"] = append(ctx.handlers["edited_message"], handler)
}

func (ctx *Client) OnChannelPost(handler func(update types.Message)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]interface{})
	}
	ctx.handlers["channel_post"] = append(ctx.handlers["channel_post"], handler)
}

func (ctx *Client) OnEditedChannelPost(handler func(update types.Message)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]interface{})
	}
	ctx.handlers["edited_channel_post"] = append(ctx.handlers["edited_channel_post"], handler)
}

func (ctx *Client) OnInlineQuery(handler func(update types.InlineQuery)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]interface{})
	}
	ctx.handlers["inline_query"] = append(ctx.handlers["inline_query"], handler)
}

func (ctx *Client) OnChosenInlineResult(handler func(update types.ChosenInlineResult)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]interface{})
	}
	ctx.handlers["chosen_inline_result"] = append(ctx.handlers["chosen_inline_result"], handler)
}

func (ctx *Client) OnCallbackQuery(handler func(update types.CallbackQuery)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]interface{})
	}
	ctx.handlers["callback_query"] = append(ctx.handlers["callback_query"], handler)
}

func (ctx *Client) OnShippingQuery(handler func(update types.ShippingQuery)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]interface{})
	}
	ctx.handlers["shipping_query"] = append(ctx.handlers["shipping_query"], handler)
}

func (ctx *Client) OnPreCheckoutQuery(handler func(update types.PreCheckoutQuery)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]interface{})
	}
	ctx.handlers["pre_checkout_query"] = append(ctx.handlers["pre_checkout_query"], handler)
}

func (ctx *Client) OnPoll(handler func(update types.Poll)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]interface{})
	}
	ctx.handlers["poll"] = append(ctx.handlers["poll"], handler)
}

func (ctx *Client) OnPollAnswer(handler func(update types.PollAnswer)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]interface{})
	}
	ctx.handlers["poll_answer"] = append(ctx.handlers["poll_answer"], handler)
}

func (ctx *Client) OnMyChatMember(handler func(update types.ChatMemberUpdated)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]interface{})
	}
	ctx.handlers["my_chat_member"] = append(ctx.handlers["my_chat_member"], handler)
}

func (ctx *Client) OnChatMember(handler func(update types.ChatMemberUpdated)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]interface{})
	}
	ctx.handlers["chat_member"] = append(ctx.handlers["chat_member"], handler)
}

func (ctx *Client) OnChatJoinRequest(handler func(update types.ChatJoinRequest)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]interface{})
	}
	ctx.handlers["chat_join_request"] = append(ctx.handlers["chat_join_request"], handler)
}

