package gobotapi

import (
	"errors"
	"fmt"
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
	"github.com/Squirrel-Network/gobotapi/utils"
	"time"
)

func (ctx *Client) Invoke(method rawTypes.Method) (*rawTypes.Result, error) {
	if !ctx.isRunning {
		return nil, errors.New("bot is not started")
	}
	files := method.Files()
	form, err := utils.GetForm(method)
	if err != nil {
		return nil, err
	}
	rawResult, err := ctx.executeRequest(
		fmt.Sprintf(
			"%sbot%s/%s",
			ctx.apiURL,
			ctx.Token,
			method.MethodName(),
		),
		"POST",
		form,
		files,
		method.ProgressCallable(),
	)
	res, err := utils.ParseResult(rawResult, err, method)
	if res.Kind == types.TypeErrorMessage && res.Error.Parameters != nil {
		retryAfter := res.Error.Parameters.RetryAfter
		if retryAfter > 0 && retryAfter <= ctx.SleepThreshold {
			time.Sleep(time.Duration(retryAfter) * time.Second)
			return ctx.Invoke(method)
		}
	}
	return res, err
}
