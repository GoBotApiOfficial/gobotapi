package main

import (
	"github.com/Squirrel-Network/gobotapi"
	"time"
)

func main() {
	client := gobotapi.Client{
		BotApiConfig: gobotapi.Config{
			HostName: "localhost", // Your bot API hostname
			Port:     8080,        // Your bot API port
			Https:    false,       // Your bot API protocol
		},
		Token:          "YOUR_TOKEN",     // Your bot token
		PollingTimeout: time.Second * 10, // Timeout for polling
	}

	// Start and idle the bot
	client.Run()
}
