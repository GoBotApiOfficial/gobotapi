// Code AutoGenerated; DO NOT EDIT.

package gobotapi

import "github.com/GoBotApiOfficial/gobotapi/types"

func (ctx *BasicClient) handleUpdate(user *types.User, token string, update types.Update) {
	client := &Client{
		Token:       token,
		BasicClient: ctx,
		me:          user,
	}
	for _, x0 := range ctx.handlers["raw"] {
		ctx.concurrencyManager.Enqueue(func(x ...any) {
			x[0].(func(*Client, types.Update))(client, update)
		}, x0)
	}
	if update.Message != nil {
		for _, x0 := range ctx.handlers["message"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.Message))(client, *update.Message)
			}, x0)
		}
	}
	if update.EditedMessage != nil {
		for _, x0 := range ctx.handlers["edited_message"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.Message))(client, *update.EditedMessage)
			}, x0)
		}
	}
	if update.ChannelPost != nil {
		for _, x0 := range ctx.handlers["channel_post"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.Message))(client, *update.ChannelPost)
			}, x0)
		}
	}
	if update.EditedChannelPost != nil {
		for _, x0 := range ctx.handlers["edited_channel_post"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.Message))(client, *update.EditedChannelPost)
			}, x0)
		}
	}
	if update.BusinessConnection != nil {
		for _, x0 := range ctx.handlers["business_connection"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.BusinessConnection))(client, *update.BusinessConnection)
			}, x0)
		}
	}
	if update.BusinessMessage != nil {
		for _, x0 := range ctx.handlers["business_message"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.Message))(client, *update.BusinessMessage)
			}, x0)
		}
	}
	if update.EditedBusinessMessage != nil {
		for _, x0 := range ctx.handlers["edited_business_message"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.Message))(client, *update.EditedBusinessMessage)
			}, x0)
		}
	}
	if update.DeletedBusinessMessages != nil {
		for _, x0 := range ctx.handlers["deleted_business_messages"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.BusinessMessagesDeleted))(client, *update.DeletedBusinessMessages)
			}, x0)
		}
	}
	if update.MessageReaction != nil {
		for _, x0 := range ctx.handlers["message_reaction"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.MessageReactionUpdated))(client, *update.MessageReaction)
			}, x0)
		}
	}
	if update.MessageReactionCount != nil {
		for _, x0 := range ctx.handlers["message_reaction_count"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.MessageReactionCountUpdated))(client, *update.MessageReactionCount)
			}, x0)
		}
	}
	if update.InlineQuery != nil {
		for _, x0 := range ctx.handlers["inline_query"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.InlineQuery))(client, *update.InlineQuery)
			}, x0)
		}
	}
	if update.ChosenInlineResult != nil {
		for _, x0 := range ctx.handlers["chosen_inline_result"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.ChosenInlineResult))(client, *update.ChosenInlineResult)
			}, x0)
		}
	}
	if update.CallbackQuery != nil {
		for _, x0 := range ctx.handlers["callback_query"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.CallbackQuery))(client, *update.CallbackQuery)
			}, x0)
		}
	}
	if update.ShippingQuery != nil {
		for _, x0 := range ctx.handlers["shipping_query"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.ShippingQuery))(client, *update.ShippingQuery)
			}, x0)
		}
	}
	if update.PreCheckoutQuery != nil {
		for _, x0 := range ctx.handlers["pre_checkout_query"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.PreCheckoutQuery))(client, *update.PreCheckoutQuery)
			}, x0)
		}
	}
	if update.PurchasedPaidMedia != nil {
		for _, x0 := range ctx.handlers["purchased_paid_media"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.PaidMediaPurchased))(client, *update.PurchasedPaidMedia)
			}, x0)
		}
	}
	if update.Poll != nil {
		for _, x0 := range ctx.handlers["poll"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.Poll))(client, *update.Poll)
			}, x0)
		}
	}
	if update.PollAnswer != nil {
		for _, x0 := range ctx.handlers["poll_answer"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.PollAnswer))(client, *update.PollAnswer)
			}, x0)
		}
	}
	if update.MyChatMember != nil {
		for _, x0 := range ctx.handlers["my_chat_member"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.ChatMemberUpdated))(client, *update.MyChatMember)
			}, x0)
		}
	}
	if update.ChatMember != nil {
		for _, x0 := range ctx.handlers["chat_member"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.ChatMemberUpdated))(client, *update.ChatMember)
			}, x0)
		}
	}
	if update.ChatJoinRequest != nil {
		for _, x0 := range ctx.handlers["chat_join_request"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.ChatJoinRequest))(client, *update.ChatJoinRequest)
			}, x0)
		}
	}
	if update.ChatBoost != nil {
		for _, x0 := range ctx.handlers["chat_boost"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.ChatBoostUpdated))(client, *update.ChatBoost)
			}, x0)
		}
	}
	if update.RemovedChatBoost != nil {
		for _, x0 := range ctx.handlers["removed_chat_boost"] {
			ctx.concurrencyManager.Enqueue(func(x ...any) {
				x[0].(func(*Client, types.ChatBoostRemoved))(client, *update.RemovedChatBoost)
			}, x0)
		}
	}
}
