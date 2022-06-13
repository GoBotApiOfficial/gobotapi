package main

import (
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/parser"
	"github.com/Squirrel-Network/gobotapi/types"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN_HERE")
	// Set an event handler for the "Message" event
	client.OnMessage(func(client *gobotapi.Client, update types.Message) {
		_, _ = client.Invoke(&methods.SendMessage{
			ChatID:    update.Chat.ID,
			Text:      parser.Parse(update.Text, update.Entities),
			ParseMode: "HTML",
		})
	})
	// Start and idle the bot
	_ = client.Run()
}
