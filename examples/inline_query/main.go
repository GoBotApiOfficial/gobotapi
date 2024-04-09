package main

import (
	"github.com/GoBotApiOfficial/gobotapi"
	"github.com/GoBotApiOfficial/gobotapi/methods"
	"github.com/GoBotApiOfficial/gobotapi/types"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN")
	// Add listener when a user uses the bot inline
	client.OnInlineQuery(func(client *gobotapi.Client, update types.InlineQuery) {
		_, err := client.Invoke(&methods.AnswerInlineQuery{
			InlineQueryID: update.ID,
			Results: []types.InlineQueryResult{
				types.InlineQueryResultArticle{
					ID: update.ID,
					InputMessageContent: types.InputTextMessageContent{
						MessageText: "Firefox is just the chinese name (火狐) of Red Panda.\n\nKnow more about Red Panda: https://en.wikipedia.org/wiki/Red_Panda",
					},
					Title:           "Firefox is a Red Panda!",
					Description:     "Read More!",
					ThumbnailURL:    "https://upload.wikimedia.org/wikipedia/commons/f/ff/Mozilla_Firefox_logo_2013.png",
					ThumbnailHeight: 128,
					ThumbnailWidth:  128,
				},
			},
		})
		// Check if there is an error
		if err != nil {
			panic(err)
		}
	})
	// Start and idle the bot
	_ = client.Run()
}
