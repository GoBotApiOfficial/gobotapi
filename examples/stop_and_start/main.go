package main

import (
	"fmt"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/types"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN")
	// Add listener to receive only stop commands with the default alias
	client.OnCommand("stop", nil, func(client *gobotapi.Client, update types.Message) {
		fmt.Println(fmt.Sprintf("Stop command received from %d", update.Chat.ID))
		// Stop the bot and close all the connections
		client.Stop()
	})
	// Start and idle the bot
	client.Run()
}
