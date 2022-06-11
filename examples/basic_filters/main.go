package main

import (
	"fmt"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/filters"
	"github.com/Squirrel-Network/gobotapi/types"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN")

	// A listener without filters
	client.OnMessage(func(client *gobotapi.Client, update types.Message) {
		fmt.Println("Message:", update.Text)
	})

	// A listener with filters (only messages from chat with id = -100123456789)
	client.OnMessage(filters.Filter(func(client *gobotapi.Client, message types.Message) {
		fmt.Println("Message Received from", message.Chat.ID, " with text ", message.Text)
	}, filters.ChatID(-100123456789)))

	// Logical Operators (only messages from chat with id = -100123456789 and the user is 123456789)
	// Supported operators: And(), Or(), Not()
	client.OnMessage(filters.Filter(func(client *gobotapi.Client, message types.Message) {
		fmt.Println("Message Received from", message.Chat.ID, " with text ", message.Text)
	}, filters.And(filters.ChatID(-100123456789), filters.UserID(123456789))))

	// Multiple chats or users filters (only messages from -100123456789 or -100987654321)
	client.OnMessage(filters.Filter(func(client *gobotapi.Client, message types.Message) {
		fmt.Println("Message Received Contained filter", message.Chat.ID, " with text ", message.Text)
	}, filters.ChatID(-100123456789, -100987654321)))

	// Start and idle the bot
	_ = client.Run()
}
