package main

import (
	"github.com/GoBotApiOfficial/gobotapi"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN")
	// Start
	_ = client.Start()
	// Idle
	gobotapi.Idle()
}
