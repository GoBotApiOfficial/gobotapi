// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"reflect"
)

// CreateChatInviteLink Use this method to create an additional invite link for a chat
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights
// The link can be revoked using the method revokeChatInviteLink
// Returns the new invite link as ChatInviteLink object.
type CreateChatInviteLink struct {
	ChatID             any    `json:"chat_id"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
	ExpireDate         int64  `json:"expire_date,omitempty"`
	MemberLimit        int    `json:"member_limit,omitempty"`
	Name               string `json:"name,omitempty"`
}

func (entity *CreateChatInviteLink) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *CreateChatInviteLink) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity CreateChatInviteLink) MarshalJSON() ([]byte, error) {
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
	type x0 CreateChatInviteLink
	return json.Marshal((x0)(entity))
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
	result := rawTypes.Result{
		Kind:   types.TypeChatInviteLink,
		Result: x1.Result,
	}
	return &result, nil
}
