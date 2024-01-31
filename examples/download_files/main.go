package main

import (
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi"
	"github.com/GoBotApiOfficial/gobotapi/types"
	"github.com/GoBotApiOfficial/gobotapi/utils"
	"log"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN_HERE")
	// Start the bot
	client.Start()
	// Download a file from FileID and save it to outputPath (without the progress)
	err := client.DownloadFile("FILE_ID", "PATH_TO_FILE", nil)
	if err != nil {
		// Handle error
		panic(err)
	}
	// Download a file from FileID and return it as []byte (without the progress)
	bytes, err := client.DownloadBytes("FILE_ID", nil)
	if err != nil {
		// Handle error
		panic(err)
	}
	log.Println(bytes)
	// Download a file from Message and save it to outputPath
	client.OnMessage(func(client *gobotapi.Client, message types.Message) {
		// Check if the message contains any types of file
		if utils.ContainsFiles(message) {
			// Download the file (without the progress)
			err := client.DownloadMedia(message, "PATH_TO_FILE", nil)
			if err != nil {
				// Handle error
				panic(err)
			}
		}
	})
	// Download a file from FileID and save it to outputPath (with the progress)
	_ = client.DownloadFile("FILE_ID", "PATH_TO_FILE", func(downloadedBytes int64, totalBytes int64) {
		// Handle the progress
		log.Println(fmt.Sprintf("%d/%d", downloadedBytes, totalBytes))
	})
	// Gracefully stop the bot
	client.Stop()
}
