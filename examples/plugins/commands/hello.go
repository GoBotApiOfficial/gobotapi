package commands

import (
	"github.com/GoBotApiOfficial/gobotapi"
	"github.com/GoBotApiOfficial/gobotapi/methods"
	"github.com/GoBotApiOfficial/gobotapi/types"
)

func Hello(client *gobotapi.Client, update types.Message) {
	client.Invoke(&methods.SendMessage{
		ChatID: update.Chat.ID,
		Text:   "Bye!",
	})
}
