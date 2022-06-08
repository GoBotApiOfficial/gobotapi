package commands

import (
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
)

func Hello(client *gobotapi.Client, update types.Message) {
	client.Invoke(&methods.SendMessage{
		ChatID: update.Chat.ID,
		Text:   "Bye!",
	})
}
