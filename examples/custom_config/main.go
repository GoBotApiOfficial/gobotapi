package main

import (
	"github.com/Squirrel-Network/gobotapi"
	"time"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN") // A polling client, you can use also a webhook client
	client.SetupBotAPI(
		"localhost", // Your bot API hostname
		8080,        // Your bot API port
		false,       // Your bot API protocol
	)
	client.PollingTimeout = time.Second * 10            // Timeout for polling
	client.NoUpdates = true                             // If true, the bot will not receive updates
	client.DownloadRefreshRate = time.Millisecond * 200 // Rate of download refresh time (in milliseconds)
	// Start and idle the bot
	_ = client.Run()
}
