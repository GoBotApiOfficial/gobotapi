package gobotapi

import (
	"fmt"
	"github.com/Squirrel-Network/gobotapi/logger"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
	"os"
)

func NewClient(token string) *PollingClient {
	return &PollingClient{
		Client: &Client{
			BasicClient: &BasicClient{
				AntiFloodData:  make(map[int64]*rawTypes.AntiFloodData),
				SleepThreshold: 10,
				MaxGoRoutines:  -1,
				LoggerColorful: true,
				LoggingLevel:   logger.Warn,
			},
			Token: token,
		},
	}
}

func NewWebhook(hostname string, port int) *WebhookClient {
	return &WebhookClient{
		BasicClient: &BasicClient{
			AntiFloodData:  make(map[int64]*rawTypes.AntiFloodData),
			SleepThreshold: 10,
			MaxGoRoutines:  -1,
			LoggerColorful: true,
			LoggingLevel:   logger.Warn,
		},
		WebhookConfig: &WebhookConfig{
			Config: ServerConfig{
				HostName: hostname,
				Port:     port,
			},
		},
	}
}

func NewWebhookWithCert(hostname string, port int, certFile, keyFile string) (*WebhookClient, error) {
	if _, ok := os.Stat(certFile); ok != nil {
		return nil, fmt.Errorf("cert file not found")
	}
	if _, ok := os.Stat(keyFile); ok != nil {
		return nil, fmt.Errorf("key file not found")
	}
	return &WebhookClient{
		BasicClient: &BasicClient{
			AntiFloodData:  make(map[int64]*rawTypes.AntiFloodData),
			SleepThreshold: 10,
			MaxGoRoutines:  -1,
			LoggerColorful: true,
			LoggingLevel:   logger.Warn,
		},
		WebhookConfig: &WebhookConfig{
			Config: ServerConfig{
				HostName: hostname,
				Port:     port,
			},
			CertificateFile: certFile,
			KeyFile:         keyFile,
		},
	}, nil
}
