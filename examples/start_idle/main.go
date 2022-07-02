package main

import (
	"github.com/Squirrel-Network/gobotapi"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN")
	// Start
	_ = client.Start()
	// Idle
	gobotapi.Idle()
}
