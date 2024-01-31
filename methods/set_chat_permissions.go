// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// SetChatPermissions Use this method to set default chat permissions for all members
// The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members administrator rights
// Returns True on success.
type SetChatPermissions struct {
	ChatID                        any                   `json:"chat_id"`
	Permissions                   types.ChatPermissions `json:"permissions"`
	UseIndependentChatPermissions bool                  `json:"use_independent_chat_permissions,omitempty"`
}

func (entity *SetChatPermissions) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SetChatPermissions) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity SetChatPermissions) MarshalJSON() ([]byte, error) {
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 SetChatPermissions
	return json.Marshal((x0)(entity))
}

func (SetChatPermissions) MethodName() string {
	return "setChatPermissions"
}

func (SetChatPermissions) ParseResult(response []byte) (*rawTypes.Result, error) {
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
