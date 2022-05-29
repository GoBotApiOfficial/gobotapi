package gobotapi

import (
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
	"time"
)

type Client struct {
	PollingTimeout  time.Duration
	BotApiConfig    Config
	Token           string
	NoUpdates       bool
	apiURL          string
	isStarted       bool
	handlers        map[string][]interface{}
	botID           int64
	botUsername     string
	lastUpdateID    int
	waitStart       chan bool
	requestsContext []rawTypes.CancelableContext
}
