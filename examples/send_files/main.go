package main

import (
	"fmt"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN_HERE")
	// Start the bot
	client.Start()
	// Send a file from local folder
	client.Invoke(&methods.SendDocument{
		ChatID:   -100123456789,
		Document: types.InputFile("file.txt"),
	})
	// Send a file from a URL
	client.Invoke(&methods.SendDocument{
		ChatID:   -100123456789,
		Document: types.InputURL("https://upload.wikimedia.org/wikipedia/commons/1/16/Fox_-_British_Wildlife_Centre_%2817429406401%29.jpg"),
	})
	// Send a file from ID
	client.Invoke(&methods.SendDocument{
		ChatID:   -100123456789,
		Document: types.InputURL("INPUT_FILE_ID"),
	})
	// Send a file from Local BotAPI File
	client.Invoke(&methods.SendDocument{
		ChatID:   -100123456789,
		Document: types.InputURL("file://INPUT_FILE_PATH"),
	})
	// Send multiple files
	client.Invoke(&methods.SendMediaGroup{
		ChatID: -100123456789,
		Media: []types.InputMedia{
			&types.InputMediaDocument{
				Media: types.InputBytes{
					Name: "file.txt",
					Data: []byte("Hello World!"),
				},
			},
			&types.InputMediaDocument{
				Media: types.InputBytes{
					Name: "hidden_fox.txt",
					Data: []byte("Oh no, another fox!"),
				},
			},
		},
	})
	// Send a file from Local BotAPI File (with the progress)
	client.Invoke(&methods.SendDocument{
		ChatID:   -100123456789,
		Document: types.InputURL("file://INPUT_FILE_PATH"),
		Progress: func(current, total int64) {
			fmt.Printf("%d/%d\n", current, total)
		},
	})
	// Gracefully stop the bot
	client.Stop()
}
