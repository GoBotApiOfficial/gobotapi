package main

import (
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/logger"
	"log"
	"os"
	"time"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN") // A polling client, you can use also a webhook client
	client.SetupBotAPI(
		"localhost", // Your bot API hostname
		8080,        // Your bot API port
		false,       // Your bot API protocol
	)
	client.PollingTimeout = time.Second * 10                    // Timeout for polling
	client.NoUpdates = true                                     // If true, the bot will not receive updates
	client.DownloadRefreshRate = time.Millisecond * 200         // Rate of download refresh time (in milliseconds)
	client.AllowedUpdates = []string{"message"}                 // Allowed updates
	client.Beta = true                                          // If true, the bot will use the beta API
	client.SleepThreshold = 10                                  // Sleep threshold, by default is 10
	client.LoggingLevel = logger.Debug                          // Logging level
	client.LoggerColorful = true                                // If true, the logger will use colorful output
	client.LoggerWriter = log.New(os.Stdout, "", log.LstdFlags) // If you want to use a custom writer, you can set it here
	client.MaxGoRoutines = 400                                  // Max number of go routines that the bot can use
	// Start and idle the bot
	_ = client.Run()
}
