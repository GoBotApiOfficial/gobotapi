package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

type GetChat struct {
	ChatId int64 `json:"chat_id"`
}

func (entity *GetChat) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (GetChat) MethodName() string {
	return "getChat"
}

func (GetChat) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.Chat `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeChat,
		Result: x1.Result,
	}
	return &result, nil
}
