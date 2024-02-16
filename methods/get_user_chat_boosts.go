// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"gobotapi/types"
	rawTypes "gobotapi/types/raw"
)

// GetUserChatBoosts Use this method to get the list of boosts added to a chat by a user
// Requires administrator rights in the chat
// Returns a UserChatBoosts object.
type GetUserChatBoosts struct {
	ChatID any   `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

func (entity *GetUserChatBoosts) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *GetUserChatBoosts) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity GetUserChatBoosts) MarshalJSON() ([]byte, error) {
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 GetUserChatBoosts
	return json.Marshal((x0)(entity))
}

func (GetUserChatBoosts) MethodName() string {
	return "getUserChatBoosts"
}

func (GetUserChatBoosts) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.UserChatBoosts `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeUserChatBoosts,
		Result: x1.Result,
	}
	return &result, nil
}
