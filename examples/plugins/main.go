package main

import (
	"github.com/Squirrel-Network/gobotapi"
	"main/commands"
)

var AliasList = []string{"/", ".", "!"}

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN")

	client.OnCommand("start", AliasList, commands.Start)
	client.OnCommand("hello", AliasList, commands.Hello)
	client.Run()
}
