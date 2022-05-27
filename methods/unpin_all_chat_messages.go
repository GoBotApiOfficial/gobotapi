package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

type UnpinAllChatMessages struct {
	ChatId int64 `json:"chat_id"`
}

func (entity *UnpinAllChatMessages) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (UnpinAllChatMessages) MethodName() string {
	return "unpinAllChatMessages"
}

func (UnpinAllChatMessages) ParseResult(response []byte) (*rawTypes.Result, error) {
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
