// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"reflect"
)

// DeleteForumTopic Use this method to delete a forum topic along with all its messages in a forum supergroup chat
// The bot must be an administrator in the chat for this to work and must have the can_delete_messages administrator rights
// Returns True on success.
type DeleteForumTopic struct {
	ChatID          any   `json:"chat_id"`
	MessageThreadID int64 `json:"message_thread_id"`
}

func (entity *DeleteForumTopic) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *DeleteForumTopic) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity DeleteForumTopic) MarshalJSON() ([]byte, error) {
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
	type x0 DeleteForumTopic
	return json.Marshal((x0)(entity))
}

func (DeleteForumTopic) MethodName() string {
	return "deleteForumTopic"
}

func (DeleteForumTopic) ParseResult(response []byte) (*rawTypes.Result, error) {
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
