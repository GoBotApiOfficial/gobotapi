// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"gobotapi/types"
	rawTypes "gobotapi/types/raw"
)

// UnhideGeneralForumTopic Use this method to unhide the 'General' topic in a forum supergroup chat
// The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights
// Returns True on success.
type UnhideGeneralForumTopic struct {
	ChatID any `json:"chat_id"`
}

func (entity *UnhideGeneralForumTopic) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *UnhideGeneralForumTopic) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity UnhideGeneralForumTopic) MarshalJSON() ([]byte, error) {
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 UnhideGeneralForumTopic
	return json.Marshal((x0)(entity))
}

func (UnhideGeneralForumTopic) MethodName() string {
	return "unhideGeneralForumTopic"
}

func (UnhideGeneralForumTopic) ParseResult(response []byte) (*rawTypes.Result, error) {
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
