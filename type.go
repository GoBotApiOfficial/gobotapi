package gobotapi

import (
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
	"net/http"
	"time"
)

type Client struct {
	PollingTimeout      time.Duration
	BotApiConfig        Config
	Token               string
	NoUpdates           bool
	DownloadRefreshRate time.Duration
	me                  *types.User
	apiURL              string
	isRunning           bool
	client              *http.Client
	handlers            map[string][]any
	lastUpdateID        int
	waitStart           chan bool
	requestsContext     []rawTypes.CancelableContext
}
