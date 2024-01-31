package main

import (
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi"
	"github.com/GoBotApiOfficial/gobotapi/filters"
	"github.com/GoBotApiOfficial/gobotapi/types"
	cFilters "main/filters"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN")
	w := cFilters.FilterWrapper(client)

	// Custom filters (only messages from admin users)
	client.OnMessage(filters.Filter(func(ctx *gobotapi.Client, message types.Message) {
		fmt.Println("Message Received from an admin:", message.Text)
	}, w.IsAdmin()))

	// Start and idle the bot
	_ = client.Run()
}
