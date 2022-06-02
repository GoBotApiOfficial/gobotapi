package main

import (
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/types"
	"log"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN_HERE")
	// Start the bot
	client.Start()
	// Download a file from FileID and save it to outputPath
	err := client.DownloadFile("FILE_ID", "PATH_TO_FILE")
	if err != nil {
		// Handle error
		panic(err)
	}
	// Download a file from FileID and return it as []byte
	bytes, err := client.DownloadBytes("FILE_ID")
	if err != nil {
		// Handle error
		panic(err)
	}
	log.Println(bytes)
	// Download a file from Message and save it to outputPath
	client.OnMessage(func(client gobotapi.Client, message types.Message) {
		// Check if the message contains any types of file
		if gobotapi.ContainsFiles(message) {
			// Download the file
			err := client.DownloadMedia(message, "PATH_TO_FILE")
			if err != nil {
				// Handle error
				panic(err)
			}
		}
	})
	// Gracefully stop the bot
	client.Stop()
}
