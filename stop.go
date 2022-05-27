package gobotapi

func (ctx *Client) Stop() {
	if ctx.isStarted {
		ctx.isStarted = false
		for _, cancelable := range ctx.requestsContext {
			cancelable.Cancel()
		}
	}
}
