package gobotapi

import (
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
	"time"
)

type Client struct {
	PollingTimeout  time.Duration
	BotApiConfig    Config
	Token           string
	apiUrl          string
	isStarted       bool
	handlers        map[string][]interface{}
	botId           int64
	botUsername     string
	lastUpdateId    int
	waitStart       chan bool
	requestsContext []rawTypes.CancelableContext
}
