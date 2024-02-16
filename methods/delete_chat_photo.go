// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"gobotapi/types"
	rawTypes "gobotapi/types/raw"
)

// DeleteChatPhoto Use this method to delete a chat photo
// Photos can't be changed for private chats
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights
// Returns True on success.
type DeleteChatPhoto struct {
	ChatID any `json:"chat_id"`
}

func (entity *DeleteChatPhoto) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *DeleteChatPhoto) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity DeleteChatPhoto) MarshalJSON() ([]byte, error) {
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 DeleteChatPhoto
	return json.Marshal((x0)(entity))
}

func (DeleteChatPhoto) MethodName() string {
	return "deleteChatPhoto"
}

func (DeleteChatPhoto) ParseResult(response []byte) (*rawTypes.Result, error) {
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
