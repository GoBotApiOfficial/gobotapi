package gobotapi

import (
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
	"net/http"
	"time"
)

type BasicClient struct {
	BotApiConfig        Config
	NoUpdates           bool
	DownloadRefreshRate time.Duration
	AllowedUpdates      []string
	Beta                bool
	SleepThreshold      int
	apiURL              string
	cloningURL          string
	isRunning           bool
	client              *http.Client
	handlers            map[string][]any
	requestsContext     []rawTypes.CancelableContext
}

type Client struct {
	*BasicClient
	me    *types.User
	Token string
}

type PollingClient struct {
	*Client
	PollingTimeout time.Duration
	lastUpdateID   int
}

type WebhookClient struct {
	*BasicClient
	WebhookConfig *WebhookConfig
	clients       map[string]*types.User
}
