package main

import (
	"fmt"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/filters"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN")
	// Add listener to receive only start commands with alias "/", ";", "."
	client.OnAnyMessageEvent(filters.Filter(func(client *gobotapi.Client, update types.Message) {
		_, err := client.Invoke(&methods.SendMessage{
			ChatID: update.Chat.ID,
			Text:   "Hello, I'm a bot!",
		})
		// Check if there is an error
		if err != nil {
			panic(err)
		}
	}, filters.Command("start", "/", ";", ".")))
	// Add listener to receive only stop commands with the default alias
	client.OnAnyMessageEvent(filters.Filter(func(client *gobotapi.Client, update types.Message) {
		fmt.Println(fmt.Sprintf("Stop command received from %d", update.Chat.ID))
	}, filters.Command("stop")))
	// Add listener to receive only commands a list of commands
	client.OnAnyMessageEvent(filters.Filter(func(client *gobotapi.Client, update types.Message) {
		fmt.Println(fmt.Sprintf("Command received from %d", update.Chat.ID))
	}, filters.Commands([]string{"help", "info", "about"})))
	// Start and idle the bot
	_ = client.Run()
}
