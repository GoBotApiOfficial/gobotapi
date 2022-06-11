package main

import (
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/types"
	"log"
)

func main() {
	// Webhook client doesn't implement the Invoke method, only from events are available.
	// Set up a proxy to NGINX or another web server.
	// The ip will be localhost:4502.
	client := gobotapi.NewWebhook("example.com", 4502)
	// Start the client
	client.Start()
	// Set an event handler for the "Message" event
	client.OnMessage(func(client *gobotapi.Client, update types.Message) {
		log.Printf("Message: %+v", update)
	})
	// The cloning of the client works with the following way:
	// client.Clone(BotToken, Drop Updates, Max Connections)
	// The webhook url will be: https://example.com/?token=BotToken
	err := client.Clone("YOUR_TOKEN", false, 100)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
