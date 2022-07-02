package gobotapi

func (ctx *WebhookClient) Run() error {
	err := ctx.Start()
	if err == nil {
		Idle()
	}
	return err
}

func (ctx *PollingClient) Run() error {
	err := ctx.Start()
	if err == nil {
		Idle()
	}
	return err
}
