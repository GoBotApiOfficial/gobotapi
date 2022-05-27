package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

type CreateChatInviteLink struct {
	ChatId int64 `json:"chat_id"`
	CreatesJoinRequest bool `json:"creates_join_request,omitempty"`
	ExpireDate int64 `json:"expire_date,omitempty"`
	MemberLimit int `json:"member_limit,omitempty"`
	Name string `json:"name,omitempty"`
}

func (entity *CreateChatInviteLink) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (CreateChatInviteLink) MethodName() string {
	return "createChatInviteLink"
}

func (CreateChatInviteLink) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.ChatInviteLink `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeChatInviteLink,
		Result: x1.Result,
	}
	return &result, nil
}