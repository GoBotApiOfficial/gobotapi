package gobotapi

import (
	"errors"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"github.com/GoBotApiOfficial/gobotapi/utils"
	"reflect"
	"time"
)

func (ctx *Client) Invoke(method rawTypes.Method) (*rawTypes.Result, error) {
	methodName := reflect.TypeOf(method).Elem().String()
	if !ctx.isRunning {
		return nil, errors.New("bot is not started")
	}
	files := method.Files()
	form, err := utils.GetForm(method)
	if err != nil {
		ctx.logging.Error(ctx, "gobotapi.invoke:", fmt.Sprintf("Failed to get form for %s,", methodName), err.Error())
		return nil, err
	}
	ctx.logging.Debug(ctx, "gobotapi.invoke:", "Generated form:", form)
	ctx.logging.Debug(ctx, "gobotapi.invoke:", "Sent:", method)
	var rawResult []byte
	for i := 0; i < 3; i++ {
		rawResult, err = ctx.executeRequest(
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
		if utils.IsServerError(rawResult, err, method) {
			ctx.logging.Debug(ctx, "gobotapi.invoke:", "Server error, retrying...", err.Error())
			time.Sleep(time.Second)
			continue
		}
		break
	}
	res, err := utils.ParseResult(rawResult, err, method)
	if res != nil && res.Kind == types.TypeErrorMessage && res.Error.Parameters != nil {
		retryAfter := res.Error.Parameters.RetryAfter
		if retryAfter > 0 && retryAfter <= ctx.SleepThreshold {
			ctx.logging.Warn(ctx, "gobotapi.invoke:", "Waiting for", retryAfter, fmt.Sprintf("seconds before continuing (required by \"%s\")", methodName))
			time.Sleep(time.Duration(retryAfter) * time.Second)
			return ctx.Invoke(method)
		}
	}
	if err != nil {
		if res != nil {
			ctx.logging.Error(ctx, "gobotapi.invoke:", "Telegram says:", fmt.Sprintf("[%d] - %s", res.Error.Code, res.Error.Description), "for method", methodName)
		} else {
			ctx.logging.Error(ctx, "gobotapi.invoke:", err.Error(), "for method", methodName)
		}
	} else {
		if reflect.TypeOf(res.Result).Kind() == reflect.Slice {
			for i := 0; i < reflect.ValueOf(res.Result).Len(); i++ {
				iV := reflect.ValueOf(res.Result).Index(i).Interface()
				ctx.logging.Debug(ctx, "gobotapi.invoke:", "Received:", iV)
			}
		} else {
			ctx.logging.Debug(ctx, "gobotapi.invoke:", "Received:", res.Result)
		}
	}
	return res, err
}
