package gobotapi

import (
	"github.com/GoBotApiOfficial/gobotapi/utils/concurrency"
	"net/http"
	"time"
)

func (ctx *BasicClient) setup() {
	ctx.client = &http.Client{}
	ctx.apiURL = ctx.BotApiConfig.link()
	if ctx.Beta {
		ctx.apiURL += "beta/"
	}
	ctx.logging = NewLogger(ctx.LoggingLevel)
	ctx.concurrencyManager = concurrency.NewPool(ctx.MaxGoRoutines)
	if ctx.DownloadRefreshRate == 0 {
		ctx.DownloadRefreshRate = time.Millisecond * 200
	}
	ctx.isRunning = true
	if !ctx.NoNotice {
		showNotice()
	}
}
