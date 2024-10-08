// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"reflect"
)

// PromoteChatMember Use this method to promote or demote a user in a supergroup or a channel
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights
// Pass False for all boolean parameters to demote a user
// Returns True on success.
type PromoteChatMember struct {
	CanChangeInfo       bool  `json:"can_change_info,omitempty"`
	CanDeleteMessages   bool  `json:"can_delete_messages,omitempty"`
	CanDeleteStories    bool  `json:"can_delete_stories,omitempty"`
	CanEditMessages     bool  `json:"can_edit_messages,omitempty"`
	CanEditStories      bool  `json:"can_edit_stories,omitempty"`
	CanInviteUsers      bool  `json:"can_invite_users,omitempty"`
	CanManageChat       bool  `json:"can_manage_chat,omitempty"`
	CanManageTopics     bool  `json:"can_manage_topics,omitempty"`
	CanManageVideoChats bool  `json:"can_manage_video_chats,omitempty"`
	CanPinMessages      bool  `json:"can_pin_messages,omitempty"`
	CanPostMessages     bool  `json:"can_post_messages,omitempty"`
	CanPostStories      bool  `json:"can_post_stories,omitempty"`
	CanPromoteMembers   bool  `json:"can_promote_members,omitempty"`
	CanRestrictMembers  bool  `json:"can_restrict_members,omitempty"`
	ChatID              any   `json:"chat_id"`
	IsAnonymous         bool  `json:"is_anonymous,omitempty"`
	UserID              int64 `json:"user_id"`
}

func (entity *PromoteChatMember) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *PromoteChatMember) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity PromoteChatMember) MarshalJSON() ([]byte, error) {
	nilCheck := func(val any) bool {
		if val == nil {
			return true
		}
		v := reflect.ValueOf(val)
		k := v.Kind()
		switch k {
		case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
			return v.IsNil()
		default:
			return false
		}
	}
	_ = nilCheck
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 PromoteChatMember
	return json.Marshal((x0)(entity))
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
	result := rawTypes.Result{
		Kind:   types.TypeBoolean,
		Result: x1.Result,
	}
	return &result, nil
}
