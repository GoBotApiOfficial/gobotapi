package gobotapi

import (
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
	"net/http"
	"time"
)

type Client struct {
	PollingTimeout  time.Duration
	BotApiConfig    Config
	Token           string
	NoUpdates       bool
	DownloadPath    string
	apiURL          string
	isStarted       bool
	client          *http.Client
	handlers        map[string][]interface{}
	botID           int64
	botUsername     string
	lastUpdateID    int
	waitStart       chan bool
	requestsContext []rawTypes.CancelableContext
}
