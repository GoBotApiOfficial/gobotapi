package gobotapi

import (
	"fmt"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
	"github.com/Squirrel-Network/gobotapi/utils"
)

func (ctx *Client) Invoke(method rawTypes.Method) (*rawTypes.Result, error) {
	if !ctx.isStarted {
		return nil, fmt.Errorf("bot is not started")
	}
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
		form,
		method.Files(),
	)
	return utils.ParseResult(rawResult, err, method)
}
