package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

type PromoteChatMember struct {
	CanChangeInfo bool `json:"can_change_info,omitempty"`
	CanDeleteMessages bool `json:"can_delete_messages,omitempty"`
	CanEditMessages bool `json:"can_edit_messages,omitempty"`
	CanInviteUsers bool `json:"can_invite_users,omitempty"`
	CanManageChat bool `json:"can_manage_chat,omitempty"`
	CanManageVideoChats bool `json:"can_manage_video_chats,omitempty"`
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
	CanPostMessages bool `json:"can_post_messages,omitempty"`
	CanPromoteMembers bool `json:"can_promote_members,omitempty"`
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`
	ChatId int64 `json:"chat_id"`
	IsAnonymous bool `json:"is_anonymous,omitempty"`
	UserId int64 `json:"user_id"`
}

func (entity *PromoteChatMember) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (PromoteChatMember) MethodName() string {
	return "promoteChatMember"
}

func (PromoteChatMember) ParseResult(response []byte) (*rawTypes.Result, error) {
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
