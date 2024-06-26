// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// Close Use this method to close the bot instance before moving it from one local server to another
// You need to delete the webhook before calling this method to ensure that the bot isn't launched again after server restart
// The method will return error 429 in the first 10 minutes after the bot is launched
// Returns True on success
// Requires no parameters.
type Close struct{}

func (entity *Close) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *Close) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (Close) MethodName() string {
	return "close"
}

func (Close) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result bool `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeBoolean,
		Result: x1.Result,
	}
	return &result, nil
}
