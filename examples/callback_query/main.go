package main

import (
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN")
	// Add listener to start command
	client.OnCommand("start", nil, func(client *gobotapi.Client, update types.Message) {
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
	})
	// Add listener to receive callback query
	client.OnCallbackQuery(func(client *gobotapi.Client, update types.CallbackQuery) {
		if update.Data == "test" {
			client.Invoke(&methods.AnswerCallbackQuery{
				CallbackQueryID: update.ID,
				Text:            "You found a Fox 🦊!",
			})
		}
	})
	// Add listener to receive callback query with filtered data
	client.OnCallbackQueryData("test2", func(client *gobotapi.Client, update types.CallbackQuery) {
		client.Invoke(&methods.AnswerCallbackQuery{
			CallbackQueryID: update.ID,
			Text:            "You found another Fox 🦊!",
		})
	})
	// Start and idle the bot
	_ = client.Run()
}
