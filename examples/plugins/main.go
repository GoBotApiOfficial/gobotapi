package main

import (
	"github.com/GoBotApiOfficial/gobotapi"
	"github.com/GoBotApiOfficial/gobotapi/filters"
	"main/commands"
)

var AliasList = []string{"/", ".", "!"}

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN")

	// Add listener to /start commands
	client.OnAnyMessageEvent(filters.Filter(commands.Start, filters.Command("start", AliasList...)))
	// Add listener to /hello commands
	client.OnAnyMessageEvent(filters.Filter(commands.Hello, filters.Command("hello", AliasList...)))
	// Start and idle the bot
	_ = client.Run()
}
