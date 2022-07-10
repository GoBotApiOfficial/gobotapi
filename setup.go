package gobotapi

import (
	"github.com/Squirrel-Network/gobotapi/utils/concurrency"
	"net/http"
	"runtime"
	"time"
)

func (ctx *BasicClient) setup() {
	ctx.client = &http.Client{}
	ctx.apiURL = ctx.BotApiConfig.link()
	if ctx.Beta {
		ctx.apiURL += "beta/"
	}
	if ctx.MaxGoRoutines == 0 {
		ctx.MaxGoRoutines = runtime.NumCPU() * 125
	}
	ctx.concurrencyManager = concurrency.New(ctx.MaxGoRoutines)
	if ctx.DownloadRefreshRate == 0 {
		ctx.DownloadRefreshRate = time.Millisecond * 200
	}
	ctx.isRunning = true
	showNotice()
}
