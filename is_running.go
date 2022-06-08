package gobotapi

func (ctx *Client) IsRunning() bool {
	return ctx.isRunning
}
