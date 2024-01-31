package main

import (
	"fmt"
	"gobotapi"
	"gobotapi/methods"
	"gobotapi/types"
)

func main() {
	client := gobotapi.NewClient("XXXX:XXXXXXXXX")
	// Start the bot
	client.Start()
	// Invoke a SendMessage method
	result, err := client.Invoke(&methods.SendMessage{
		ChatID: 1234556789,
		Text:   "Hello World!",
	})
	// Check if there is an error
	if err != nil {
		panic(err)
	}
	// Kind is the type of the result
	if result.Kind == types.TypeMessage {
		message := result.Result.(types.Message)
		// Print the message result
		fmt.Println(message)
	}
	// Stop the bot to close all the connections
	client.Stop()
}
