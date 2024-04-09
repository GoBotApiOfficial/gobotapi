// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// GetMyShortDescription Use this method to get the current bot short description for the given user language
// Returns BotShortDescription on success.
type GetMyShortDescription struct {
	LanguageCode string `json:"language_code,omitempty"`
}

func (entity *GetMyShortDescription) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *GetMyShortDescription) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (GetMyShortDescription) MethodName() string {
	return "getMyShortDescription"
}

func (GetMyShortDescription) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.BotShortDescription `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeBotShortDescription,
		Result: x1.Result,
	}
	return &result, nil
}
