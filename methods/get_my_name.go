// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// GetMyName Use this method to get the current bot name for the given user language
// Returns BotName on success.
type GetMyName struct {
	LanguageCode string `json:"language_code,omitempty"`
}

func (entity *GetMyName) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *GetMyName) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (GetMyName) MethodName() string {
	return "getMyName"
}

func (GetMyName) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.BotName `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeBotName,
		Result: x1.Result,
	}
	return &result, nil
}
