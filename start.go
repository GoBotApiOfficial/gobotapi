package gobotapi

func (ctx *Client) Start() {
	ctx.waitStart = make(chan bool)
	go ctx.Run()
	<-ctx.waitStart
}
