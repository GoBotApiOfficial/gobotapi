package main

import (
	"github.com/Squirrel-Network/gobotapi"
)

func main() {
	client := gobotapi.Client{
		BotApiConfig: gobotapi.Config{
			HostName: "your_host", // Your host name
			Port:     443,         // Your port
			Https:    true,        // Your protocol
		},
		Token: "YOUR_TOKEN",
	}
	// Start and idle the bot
	client.Run()
}
