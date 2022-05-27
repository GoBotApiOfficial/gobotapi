package methods

import "github.com/Squirrel-Network/gobotapi/types"
import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "encoding/json"

type DeleteMyCommands struct {
	LanguageCode string `json:"language_code,omitempty"`
	Scope *types.BotCommandScope `json:"scope,omitempty"`
}

func (entity *DeleteMyCommands) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (DeleteMyCommands) MethodName() string {
	return "deleteMyCommands"
}

func (DeleteMyCommands) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result bool `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeBoolean,
		Result: x1.Result,
	}
	return &result, nil
}
