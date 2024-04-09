package gobotapi

import (
	"errors"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/methods"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

func (ctx *Client) DownloadBytes(fileId string, progress rawTypes.ProgressCallable) ([]byte, error) {
	if len(ctx.BotApiConfig.HostName) > 0 && ctx.BotApiConfig.HostName != "api.telegram.org" {
		return nil, errors.New("download is supported only on api.telegram.org")
	}
	invoke, err := ctx.Invoke(&methods.GetFile{
		FileID: fileId,
	})
	if err != nil {
		return nil, err
	}
	return ctx.executeRequest(
		fmt.Sprintf(
			"%sfile/bot%s/%s",
			ctx.apiURL,
			ctx.Token,
			invoke.Result.(types.File).FilePath,
		),
		"GET",
		nil,
		nil,
		progress,
	)
}
