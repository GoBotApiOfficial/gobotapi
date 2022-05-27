package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

type GetChatMemberCount struct {
	ChatId int64 `json:"chat_id"`
}

func (entity *GetChatMemberCount) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (GetChatMemberCount) MethodName() string {
	return "getChatMemberCount"
}

func (GetChatMemberCount) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result int `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeInteger,
		Result: x1.Result,
	}
	return &result, nil
}