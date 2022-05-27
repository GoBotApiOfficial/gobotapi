package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

type UnbanChatMember struct {
	ChatId int64 `json:"chat_id"`
	OnlyIfBanned bool `json:"only_if_banned,omitempty"`
	UserId int64 `json:"user_id"`
}

func (entity *UnbanChatMember) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (UnbanChatMember) MethodName() string {
	return "unbanChatMember"
}

func (UnbanChatMember) ParseResult(response []byte) (*rawTypes.Result, error) {
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
