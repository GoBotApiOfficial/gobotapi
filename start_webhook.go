package gobotapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
	"net/http"
)

func (ctx *WebhookClient) Start() error {
	if ctx.isRunning {
		return errors.New("client is already running")
	}
	if ctx.clients == nil {
		ctx.clients = make(map[string]*types.User)
	}
	ctx.BasicClient.cloningURL = fmt.Sprintf(
		"https://%s:%d/",
		ctx.WebhookConfig.Config.HostName,
		ctx.WebhookConfig.Config.Port,
	)
	ctx.BasicClient.setup()
	if ctx.NoUpdates {
		return nil
	}
	ctx.WebhookConfig.server = &http.Server{
		Addr: ctx.WebhookConfig.GetAddress(),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.URL.Query().Get("token")
			if r.URL.Path != "/" || r.Method != "POST" || len(token) == 0 {
				http.NotFound(w, r)
				return
			}
			var update types.Update
			err := json.NewDecoder(r.Body).Decode(&update)
			if err != nil {
				errMsg, _ := json.Marshal(map[string]string{"error": "Unable to parse update"})
				w.WriteHeader(http.StatusBadRequest)
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write(errMsg)
				return
			}
			ctx.mutex.Lock()
			if ctx.clients[token] == nil {
				ctx.logging.Info(nil, "Connecting...")
				client := &Client{
					Token:       token,
					BasicClient: ctx.BasicClient,
				}
				invoke, _ := client.Invoke(&methods.GetMe{})
				me := invoke.Result.(types.User)
				ctx.clients[token] = &me
				ctx.logging.Info(&Client{me: ctx.clients[token]}, "Connected")
			}
			tmpC := &Client{me: ctx.clients[token]}
			ctx.logging.Info(tmpC, "Received", 1, "updates")
			ctx.logging.Debug(tmpC, "Received:", update)
			ctx.handleUpdate(ctx.clients[token], token, update)
			ctx.mutex.Unlock()
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("OK"))
		}),
	}
	ctx.logging.Warn(nil, "Webhook server started at", ctx.WebhookConfig.GetAddress())
	var err error
	if len(ctx.WebhookConfig.CertificateFile) > 0 && len(ctx.WebhookConfig.KeyFile) > 0 {
		go func() {
			_ = ctx.WebhookConfig.server.ListenAndServeTLS(ctx.WebhookConfig.CertificateFile, ctx.WebhookConfig.KeyFile)
		}()
	} else {
		go func() {
			_ = ctx.WebhookConfig.server.ListenAndServe()
		}()
	}
	return err
}
