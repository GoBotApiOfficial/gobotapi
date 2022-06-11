package gobotapi

import (
	"fmt"
	"os"
)

func NewClient(token string) *PollingClient {
	return &PollingClient{
		Client: &Client{
			BasicClient: &BasicClient{},
			Token:       token,
		},
	}
}

func NewWebhook(hostname string, port int) *WebhookClient {
	return &WebhookClient{
		BasicClient: &BasicClient{},
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
		BasicClient: &BasicClient{},
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
