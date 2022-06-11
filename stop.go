package gobotapi

import "context"

func (ctx *BasicClient) Stop() {
	if ctx.isRunning {
		ctx.isRunning = false
		for _, cancelable := range ctx.requestsContext {
			cancelable.Cancel()
		}
	}
}

func (ctx *WebhookClient) Stop() {
	if ctx.WebhookConfig != nil && ctx.WebhookConfig.server != nil {
		_ = ctx.WebhookConfig.server.Shutdown(context.Background())
	}
	ctx.BasicClient.Stop()
}
