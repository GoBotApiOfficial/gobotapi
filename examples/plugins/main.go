package main

import (
	"github.com/Squirrel-Network/gobotapi"
	"main/commands"
)

var AliasList = []string{"/", ".", "!"}

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN")

	// Add listener to /start commands
	client.OnCommand("start", AliasList, commands.Start)
	// Add listener to /hello commands
	client.OnCommand("hello", AliasList, commands.Hello)
	// Start and idle the bot
	_ = client.Run()
}
