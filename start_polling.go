package gobotapi

import (
	"errors"
	"github.com/GoBotApiOfficial/gobotapi/methods"
	"github.com/GoBotApiOfficial/gobotapi/types"
	"time"
)

func (ctx *PollingClient) Start() error {
	if ctx.isRunning {
		return errors.New("client is already running")
	}
	if ctx.PollingTimeout == 0 {
		ctx.PollingTimeout = time.Second * 15
	}
	ctx.BasicClient.setup()
	if ctx.NoUpdates {
		return nil
	}
	ctx.logging.Info(nil, "Connecting...")
	res, err := ctx.Invoke(&methods.GetMe{})
	if err != nil {
		return err
	}
	me := res.Result.(types.User)
	ctx.me = &me
	ctx.logging.Info(ctx.Client, "Connected")
	go func() {
		for {
			getUpdates := &methods.GetUpdates{
				AllowedUpdates: ctx.AllowedUpdates,
				Timeout:        int(ctx.PollingTimeout.Seconds()),
				Offset:         ctx.lastUpdateID,
			}
			rawUpdates, errUpdate := ctx.Invoke(getUpdates)
			if !ctx.isRunning {
				break
			}
			if errUpdate != nil {
				time.Sleep(time.Second * 5)
				continue
			}
			updates := rawUpdates.Result.([]types.Update)
			if len(updates) > 0 {
				ctx.logging.Info(ctx.Client, "Received", len(updates), "updates")
			}
			for _, update := range updates {
				ctx.lastUpdateID = int(update.UpdateID) + 1
				ctx.handleUpdate(ctx.me, ctx.Token, update)
			}
		}
	}()
	return nil
}
