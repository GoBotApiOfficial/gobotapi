package gobotapi

import (
	"github.com/Squirrel-Network/gobotapi/logger"
	"github.com/Squirrel-Network/gobotapi/utils/concurrency"
	"log"
	"net/http"
	"os"
	"time"
)

func (ctx *BasicClient) setup() {
	ctx.client = &http.Client{}
	ctx.apiURL = ctx.BotApiConfig.link()
	if ctx.Beta {
		ctx.apiURL += "beta/"
	}
	if ctx.LoggerWriter == nil {
		ctx.LoggerWriter = log.New(os.Stdout, "", log.LstdFlags)
	}
	ctx.logging = NewLogger(ctx.LoggerWriter, logger.Config{
		Colorful: ctx.LoggerColorful,
		LogLevel: ctx.LoggingLevel,
	})
	ctx.concurrencyManager = concurrency.NewPool(ctx.MaxGoRoutines)
	if ctx.DownloadRefreshRate == 0 {
		ctx.DownloadRefreshRate = time.Millisecond * 200
	}
	ctx.isRunning = true
	showNotice()
}
