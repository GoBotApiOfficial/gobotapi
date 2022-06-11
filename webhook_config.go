package gobotapi

import (
	"fmt"
	"net/http"
)

type WebhookConfig struct {
	Config          ServerConfig
	CertificateFile string
	KeyFile         string
	server          *http.Server
}

func (ctx *WebhookConfig) GetAddress() string {
	ip := "0.0.0.0"
	if len(ctx.CertificateFile) == 0 || len(ctx.KeyFile) == 0 {
		ip = "localhost"
	}
	return fmt.Sprintf("%s:%d", ip, ctx.Config.Port)
}
