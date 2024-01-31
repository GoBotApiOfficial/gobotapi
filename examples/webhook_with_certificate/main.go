package main

import (
	"github.com/GoBotApiOfficial/gobotapi"
	"github.com/GoBotApiOfficial/gobotapi/types"
	"log"
)

func main() {
	// Webhook client doesn't implement the Invoke method, only from events are available.
	// Don't use this way if you're already using Nginx or Apache.
	client, err := gobotapi.NewWebhookWithCert("example.com", 443, "./cert.pem", "./key.pem")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	// Start the client
	client.Start()
	// Set an event handler for the "Message" event
	client.OnMessage(func(client *gobotapi.Client, update types.Message) {
		log.Printf("Message: %+v", update)
	})
	// The cloning of the client works with the following way:
	// client.Clone(BotToken, Drop Updates, Max Connections)
	// The webhook url will be: https://example.com/?token=BotToken
	err = client.Clone("YOUR_TOKEN", false, 100)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
