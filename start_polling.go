package gobotapi

import (
	"errors"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
	"log"
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
	res, err := ctx.Invoke(&methods.GetMe{})
	if err != nil {
		return err
	}
	me := res.Result.(types.User)
	ctx.me = &me
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
				log.Printf("[%d] Retrying \"getUpdates\" due to Telegram says %s", ctx.me.ID, err)
				time.Sleep(time.Second * 5)
				continue
			}
			updates := rawUpdates.Result.([]types.Update)
			for _, update := range updates {
				ctx.lastUpdateID = int(update.UpdateID) + 1
				ctx.handleUpdate(ctx.me, ctx.Token, update)
			}
		}
	}()
	return nil
}
