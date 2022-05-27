package main

import (
	"fmt"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/types"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN")
	// Add listener to receive only stop commands with the default alias
	client.OnCommand("stop", nil, func(update types.Message) {
		fmt.Println(fmt.Sprintf("Stop command received from %d", update.Chat.Id))
		// Stop the bot and close all connections
		client.Stop()
	})
	// Start and idle the bot
	client.Run()
}
