package gobotapi

func (ctx *Client) Stop() {
	if ctx.isRunning {
		ctx.isRunning = false
		for _, cancelable := range ctx.requestsContext {
			cancelable.Cancel()
		}
	}
}
