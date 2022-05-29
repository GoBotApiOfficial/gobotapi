// Code AutoGenerated; DO NOT EDIT.

package gobotapi

import (
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
	"log"
	"net/http"
	"time"
)


func (ctx *Client) Run() {
	if ctx.isStarted {
		return
	}
	ctx.client = &http.Client{}
	ctx.apiURL = ctx.BotApiConfig.link()
	if ctx.PollingTimeout == 0 {
		ctx.PollingTimeout = time.Second * 15
	}
	ctx.isStarted = true
	if ctx.waitStart != nil {
		ctx.waitStart <- true
	}
	res, err := ctx.Invoke(&methods.GetMe{})
	if err != nil {
		log.Fatal(err)
	}
	ctx.botID = res.Result.(types.User).ID
	ctx.botUsername = res.Result.(types.User).Username
	showNotice()
	if ctx.NoUpdates {
		return
	}
	for {
		getUpdates := &methods.GetUpdates {
			Timeout: int(ctx.PollingTimeout.Seconds()),
			Offset: ctx.lastUpdateID,
		}
		rawUpdates, err := ctx.Invoke(getUpdates)
		if !ctx.isStarted {
			break
		}
		if err != nil {
			log.Printf("[%d] Retrying \"getUpdates\" due to Telegram says %s", ctx.botID, err)
			time.Sleep(time.Second * 5)
			continue
		}
		updates := rawUpdates.Result.([]types.Update)
		for _, update := range updates {
			ctx.lastUpdateID = int(update.UpdateID) + 1
			for _, x0 := range ctx.handlers["raw"] {
				go x0.(func(types.Update))(update)
			}
			if update.Message != nil {
				for _, x0 := range ctx.handlers["message"] {
					go x0.(func(types.Message))(*update.Message)
				}
			}
			if update.EditedMessage != nil {
				for _, x0 := range ctx.handlers["edited_message"] {
					go x0.(func(types.Message))(*update.EditedMessage)
				}
			}
			if update.ChannelPost != nil {
				for _, x0 := range ctx.handlers["channel_post"] {
					go x0.(func(types.Message))(*update.ChannelPost)
				}
			}
			if update.EditedChannelPost != nil {
				for _, x0 := range ctx.handlers["edited_channel_post"] {
					go x0.(func(types.Message))(*update.EditedChannelPost)
				}
			}
			if update.InlineQuery != nil {
				for _, x0 := range ctx.handlers["inline_query"] {
					go x0.(func(types.InlineQuery))(*update.InlineQuery)
				}
			}
			if update.ChosenInlineResult != nil {
				for _, x0 := range ctx.handlers["chosen_inline_result"] {
					go x0.(func(types.ChosenInlineResult))(*update.ChosenInlineResult)
				}
			}
			if update.CallbackQuery != nil {
				for _, x0 := range ctx.handlers["callback_query"] {
					go x0.(func(types.CallbackQuery))(*update.CallbackQuery)
				}
			}
			if update.ShippingQuery != nil {
				for _, x0 := range ctx.handlers["shipping_query"] {
					go x0.(func(types.ShippingQuery))(*update.ShippingQuery)
				}
			}
			if update.PreCheckoutQuery != nil {
				for _, x0 := range ctx.handlers["pre_checkout_query"] {
					go x0.(func(types.PreCheckoutQuery))(*update.PreCheckoutQuery)
				}
			}
			if update.Poll != nil {
				for _, x0 := range ctx.handlers["poll"] {
					go x0.(func(types.Poll))(*update.Poll)
				}
			}
			if update.PollAnswer != nil {
				for _, x0 := range ctx.handlers["poll_answer"] {
					go x0.(func(types.PollAnswer))(*update.PollAnswer)
				}
			}
			if update.MyChatMember != nil {
				for _, x0 := range ctx.handlers["my_chat_member"] {
					go x0.(func(types.ChatMemberUpdated))(*update.MyChatMember)
				}
			}
			if update.ChatMember != nil {
				for _, x0 := range ctx.handlers["chat_member"] {
					go x0.(func(types.ChatMemberUpdated))(*update.ChatMember)
				}
			}
			if update.ChatJoinRequest != nil {
				for _, x0 := range ctx.handlers["chat_join_request"] {
					go x0.(func(types.ChatJoinRequest))(*update.ChatJoinRequest)
				}
			}
		}
	}
}
