package gobotapi

import (
	"gobotapi/logger"
	"gobotapi/types"
	rawTypes "gobotapi/types/raw"
	"gobotapi/utils/concurrency"
	"net/http"
	"sync"
	"time"
)

type BasicClient struct {
	BotApiConfig        Config
	NoUpdates           bool
	DownloadRefreshRate time.Duration
	AllowedUpdates      []string
	Beta                bool
	SleepThreshold      int
	MaxGoRoutines       int
	LoggingLevel        logger.LogLevel
	LoggerColorful      bool
	LoggerWriter        Writer
	AntiFloodData       map[int64]*rawTypes.AntiFloodData
	logging             Logger
	apiURL              string
	cloningURL          string
	isRunning           bool
	client              *http.Client
	handlers            map[string][]any
	requestsContext     []rawTypes.CancelableContext
	concurrencyManager  *concurrency.Pool
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
	mutex         sync.RWMutex
}
