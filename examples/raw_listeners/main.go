package main

import (
	"fmt"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN")
	// Add listener to messages
	client.OnMessage(func(update types.Message) {
		_, err := client.Invoke(&methods.SendMessage{
			ChatID: update.Chat.ID,
			Text:   "Hello, I'm a bot!",
		})
		// Check if there is an error
		if err != nil {
			panic(err)
		}
	})
	// Add listener to receive all updates
	client.OnRawUpdate(func(update types.Update) {
		// Print the update
		fmt.Println(update)
	})
	// Start and idle the bot
	client.Run()
}
