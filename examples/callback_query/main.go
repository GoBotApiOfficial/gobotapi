package main

import (
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/filters"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
	"github.com/Squirrel-Network/gobotapi/utils"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN")
	// Add listener to start command
	client.OnAnyMessageEvent(filters.Filter(func(client *gobotapi.Client, update types.Message) {
		client.Invoke(&methods.SendMessage{
			ChatID: update.Chat.ID,
			Text:   "Hello, I'm a bot!",
			ReplyMarkup: &types.InlineKeyboardMarkup{
				InlineKeyboard: [][]types.InlineKeyboardButton{
					{
						{
							Text:         "Click me!",
							CallbackData: "test",
						},
						{
							Text:         "Click me!",
							CallbackData: "test2",
						},
					},
				},
			},
		})
	}, filters.Command("start")))
	// Add listener to receive callback query
	client.OnCallbackQuery(func(client *gobotapi.Client, update types.CallbackQuery) {
		if update.Data == "test" {
			client.Invoke(&methods.AnswerCallbackQuery{
				CallbackQueryID: update.ID,
				Text:            "You found a Fox 🦊!",
			})
		}
	})

	// Split buttons by 2
	client.OnAnyMessageEvent(filters.Filter(func(client *gobotapi.Client, update types.Message) {
		client.Invoke(&methods.SendMessage{
			ChatID: update.Chat.ID,
			Text:   "Hello, I'm a bot!",
			ReplyMarkup: &types.InlineKeyboardMarkup{
				InlineKeyboard: utils.SplitButtons(2,
					types.InlineKeyboardButton{
						Text:         "Click me!",
						CallbackData: "test",
					},
					types.InlineKeyboardButton{
						Text:         "Click me!",
						CallbackData: "test2",
					},
					types.InlineKeyboardButton{
						Text:         "Click me!",
						CallbackData: "test3",
					},
					types.InlineKeyboardButton{
						Text:         "Click me!",
						CallbackData: "test4",
					},
				),
			},
		})
	}, filters.Command("split")))
	// Start and idle the bot
	_ = client.Run()
}
