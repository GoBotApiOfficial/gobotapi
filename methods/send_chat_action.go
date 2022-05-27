package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

type SendChatAction struct {
	Action string `json:"action"`
	ChatId int64 `json:"chat_id"`
}

func (entity *SendChatAction) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (SendChatAction) MethodName() string {
	return "sendChatAction"
}

func (SendChatAction) ParseResult(response []byte) (*rawTypes.Result, error) {
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
