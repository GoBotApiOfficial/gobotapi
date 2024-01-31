package main

import (
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi"
	"github.com/GoBotApiOfficial/gobotapi/filters"
	"github.com/GoBotApiOfficial/gobotapi/types"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN")
	// Add listener to receive only stop commands with the default alias
	client.OnAnyMessageEvent(filters.Filter(func(client *gobotapi.Client, update types.Message) {
		fmt.Println(fmt.Sprintf("Stop command received from %d", update.Chat.ID))
		// Stop the bot and close all the connections
		client.Stop()
	}, filters.Command("stop")))
	// Start and idle the bot
	_ = client.Run()
}
