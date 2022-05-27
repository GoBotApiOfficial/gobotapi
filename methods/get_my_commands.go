package methods

import "github.com/Squirrel-Network/gobotapi/types"
import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "encoding/json"

type GetMyCommands struct {
	LanguageCode string `json:"language_code,omitempty"`
	Scope *types.BotCommandScope `json:"scope,omitempty"`
}

func (entity *GetMyCommands) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (GetMyCommands) MethodName() string {
	return "getMyCommands"
}

func (GetMyCommands) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result []types.BotCommand `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeArrayOfBotCommand,
		Result: x1.Result,
	}
	return &result, nil
}
