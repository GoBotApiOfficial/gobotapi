package main

import (
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	client := gobotapi.NewClient("YOUR_TOKEN")
	// Read more about scheduling functions at https://godoc.org/github.com/robfig/cron/v3
	crontab := cron.New()
	// Repeat every minute
	crontab.AddFunc("@every 1m", func() {
		client.Invoke(&methods.SendMessage{
			ChatID: 123456789,
			Text:   "Hello World",
		})
	})
	// Repeat every day at 00:00
	crontab.AddFunc("@every 0 0 0 * * *", func() {
		client.Invoke(&methods.SendMessage{
			ChatID: 123456789,
			Text:   "Hello World",
		})
	})
	// Send message after 5 seconds
	time.AfterFunc(5*time.Second, func() {
		client.Invoke(&methods.SendMessage{
			ChatID: 123456789,
			Text:   "Hello World",
		})
	})
	// Start the cron job
	crontab.Start()
	// Start and idle the bot
	_ = client.Run()
}
