package main

import (
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN_HERE")
	// Start the bot
	client.Start()
	// Send a file from bytes
	client.Invoke(&methods.SendDocument{
		ChatId: -100123456789,
		Document: types.InputFile{
			Name:  "file.txt",
			Bytes: []byte("Hello World!"),
		},
	})
	// Send a file from a URL
	client.Invoke(&methods.SendDocument{
		ChatId:   -100123456789,
		Document: types.InputURL("https://upload.wikimedia.org/wikipedia/commons/1/16/Fox_-_British_Wildlife_Centre_%2817429406401%29.jpg"),
	})
	// Send multiple files
	client.Invoke(&methods.SendMediaGroup{
		ChatId: -100123456789,
		Media: []types.InputMedia{
			&types.InputMediaDocument{
				Media: types.InputFile{
					Name:  "file.txt",
					Bytes: []byte("Hello World!"),
				},
			},
			&types.InputMediaDocument{
				Media: types.InputFile{
					Name:  "hidden_fox.txt",
					Bytes: []byte("Oh no, another fox!"),
				},
			},
		},
	})
	// Gracefully stop the bot
	client.Stop()
}
