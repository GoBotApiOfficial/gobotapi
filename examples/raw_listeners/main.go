package main

import (
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi"
	"github.com/GoBotApiOfficial/gobotapi/methods"
	"github.com/GoBotApiOfficial/gobotapi/types"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN")
	// Add listener to messages
	client.OnMessage(func(client *gobotapi.Client, update types.Message) {
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
	client.OnRawUpdate(func(client *gobotapi.Client, update types.Update) {
		// Print the update
		fmt.Println(update)
	})
	// Start and idle the bot
	_ = client.Run()
}
