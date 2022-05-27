package methods

import "github.com/Squirrel-Network/gobotapi/types"
import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "encoding/json"

type RestrictChatMember struct {
	ChatId int64 `json:"chat_id"`
	Permissions types.ChatPermissions `json:"permissions"`
	UntilDate int64 `json:"until_date,omitempty"`
	UserId int64 `json:"user_id"`
}

func (entity *RestrictChatMember) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (RestrictChatMember) MethodName() string {
	return "restrictChatMember"
}

func (RestrictChatMember) ParseResult(response []byte) (*rawTypes.Result, error) {
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
