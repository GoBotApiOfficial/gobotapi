package gobotapi

import (
	"net/http"
	"time"
)

func (ctx *BasicClient) setup() {
	ctx.client = &http.Client{}
	ctx.apiURL = ctx.BotApiConfig.link()
	if ctx.Beta {
		ctx.apiURL += "beta/"
	}
	if ctx.DownloadRefreshRate == 0 {
		ctx.DownloadRefreshRate = time.Millisecond * 200
	}
	ctx.isRunning = true
	showNotice()
}
