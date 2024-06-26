// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// SetMyName Use this method to change the bot's name
// Returns True on success.
type SetMyName struct {
	LanguageCode string `json:"language_code,omitempty"`
	Name         string `json:"name,omitempty"`
}

func (entity *SetMyName) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SetMyName) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (SetMyName) MethodName() string {
	return "setMyName"
}

func (SetMyName) ParseResult(response []byte) (*rawTypes.Result, error) {
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
