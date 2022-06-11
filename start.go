package gobotapi

func (ctx *WebhookClient) Start() {
	ctx.waitStart = make(chan bool)
	go func() {
		_ = ctx.Run()
	}()
	<-ctx.waitStart
}

func (ctx *PollingClient) Start() {
	ctx.waitStart = make(chan bool)
	go func() {
		_ = ctx.Run()
	}()
	<-ctx.waitStart
}
