// Code AutoGenerated; DO NOT EDIT.

package gobotapi

import "github.com/GoBotApiOfficial/gobotapi/types"

func (ctx *BasicClient) OnRawUpdate(handler func(client *Client, update types.Update)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["raw"] = append(ctx.handlers["raw"], handler)
}

func (ctx *BasicClient) OnMessage(handler func(client *Client, update types.Message)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["message"] = append(ctx.handlers["message"], handler)
}

func (ctx *BasicClient) OnEditedMessage(handler func(client *Client, update types.Message)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["edited_message"] = append(ctx.handlers["edited_message"], handler)
}

func (ctx *BasicClient) OnChannelPost(handler func(client *Client, update types.Message)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["channel_post"] = append(ctx.handlers["channel_post"], handler)
}

func (ctx *BasicClient) OnEditedChannelPost(handler func(client *Client, update types.Message)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["edited_channel_post"] = append(ctx.handlers["edited_channel_post"], handler)
}

func (ctx *BasicClient) OnBusinessConnection(handler func(client *Client, update types.BusinessConnection)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["business_connection"] = append(ctx.handlers["business_connection"], handler)
}

func (ctx *BasicClient) OnBusinessMessage(handler func(client *Client, update types.Message)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["business_message"] = append(ctx.handlers["business_message"], handler)
}

func (ctx *BasicClient) OnEditedBusinessMessage(handler func(client *Client, update types.Message)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["edited_business_message"] = append(ctx.handlers["edited_business_message"], handler)
}

func (ctx *BasicClient) OnDeletedBusinessMessages(handler func(client *Client, update types.BusinessMessagesDeleted)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["deleted_business_messages"] = append(ctx.handlers["deleted_business_messages"], handler)
}

func (ctx *BasicClient) OnMessageReaction(handler func(client *Client, update types.MessageReactionUpdated)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["message_reaction"] = append(ctx.handlers["message_reaction"], handler)
}

func (ctx *BasicClient) OnMessageReactionCount(handler func(client *Client, update types.MessageReactionCountUpdated)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["message_reaction_count"] = append(ctx.handlers["message_reaction_count"], handler)
}

func (ctx *BasicClient) OnInlineQuery(handler func(client *Client, update types.InlineQuery)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["inline_query"] = append(ctx.handlers["inline_query"], handler)
}

func (ctx *BasicClient) OnChosenInlineResult(handler func(client *Client, update types.ChosenInlineResult)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["chosen_inline_result"] = append(ctx.handlers["chosen_inline_result"], handler)
}

func (ctx *BasicClient) OnCallbackQuery(handler func(client *Client, update types.CallbackQuery)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["callback_query"] = append(ctx.handlers["callback_query"], handler)
}

func (ctx *BasicClient) OnShippingQuery(handler func(client *Client, update types.ShippingQuery)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["shipping_query"] = append(ctx.handlers["shipping_query"], handler)
}

func (ctx *BasicClient) OnPreCheckoutQuery(handler func(client *Client, update types.PreCheckoutQuery)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["pre_checkout_query"] = append(ctx.handlers["pre_checkout_query"], handler)
}

func (ctx *BasicClient) OnPurchasedPaidMedia(handler func(client *Client, update types.PaidMediaPurchased)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["purchased_paid_media"] = append(ctx.handlers["purchased_paid_media"], handler)
}

func (ctx *BasicClient) OnPoll(handler func(client *Client, update types.Poll)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["poll"] = append(ctx.handlers["poll"], handler)
}

func (ctx *BasicClient) OnPollAnswer(handler func(client *Client, update types.PollAnswer)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["poll_answer"] = append(ctx.handlers["poll_answer"], handler)
}

func (ctx *BasicClient) OnMyChatMember(handler func(client *Client, update types.ChatMemberUpdated)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["my_chat_member"] = append(ctx.handlers["my_chat_member"], handler)
}

func (ctx *BasicClient) OnChatMember(handler func(client *Client, update types.ChatMemberUpdated)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["chat_member"] = append(ctx.handlers["chat_member"], handler)
}

func (ctx *BasicClient) OnChatJoinRequest(handler func(client *Client, update types.ChatJoinRequest)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["chat_join_request"] = append(ctx.handlers["chat_join_request"], handler)
}

func (ctx *BasicClient) OnChatBoost(handler func(client *Client, update types.ChatBoostUpdated)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["chat_boost"] = append(ctx.handlers["chat_boost"], handler)
}

func (ctx *BasicClient) OnRemovedChatBoost(handler func(client *Client, update types.ChatBoostRemoved)) {
	if ctx.handlers == nil {
		ctx.handlers = make(map[string][]any)
	}
	ctx.handlers["removed_chat_boost"] = append(ctx.handlers["removed_chat_boost"], handler)
}
