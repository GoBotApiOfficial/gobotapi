// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"gobotapi/types"
	rawTypes "gobotapi/types/raw"
)

// SetChatAdministratorCustomTitle Use this method to set a custom title for an administrator in a supergroup promoted by the bot
// Returns True on success.
type SetChatAdministratorCustomTitle struct {
	ChatID      any    `json:"chat_id"`
	CustomTitle string `json:"custom_title"`
	UserID      int64  `json:"user_id"`
}

func (entity *SetChatAdministratorCustomTitle) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SetChatAdministratorCustomTitle) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity SetChatAdministratorCustomTitle) MarshalJSON() ([]byte, error) {
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 SetChatAdministratorCustomTitle
	return json.Marshal((x0)(entity))
}

func (SetChatAdministratorCustomTitle) MethodName() string {
	return "setChatAdministratorCustomTitle"
}

func (SetChatAdministratorCustomTitle) ParseResult(response []byte) (*rawTypes.Result, error) {
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
