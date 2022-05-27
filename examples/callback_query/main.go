package main

import (
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN")
	// Add listener to start command
	client.OnCommand("start", nil, func(update types.Message) {
		client.Invoke(&methods.SendMessage{
			ChatId: update.Chat.Id,
			Text:   "Hello, I'm a bot!",
			ReplyMarkup: &types.InlineKeyboardMarkup{
				InlineKeyboard: [][]types.InlineKeyboardButton{
					{
						{
							Text:         "Click me!",
							CallbackData: "test",
						},
					},
				},
			},
		})
	})
	// Add listener to receive callback query
	client.OnCallbackQuery(func(update types.CallbackQuery) {
		if update.Data == "test" {
			client.Invoke(&methods.AnswerCallbackQuery{
				CallbackQueryId: update.Id,
				Text:            "You found an Fox 🦊!",
			})
		}
	})
	// Start and idle the bot
	client.Run()
}
