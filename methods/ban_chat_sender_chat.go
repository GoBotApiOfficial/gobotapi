// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"reflect"
)

// BanChatSenderChat Use this method to ban a channel chat in a supergroup or a channel
// Until the chat is unbanned, the owner of the banned chat won't be able to send messages on behalf of any of their channels
// The bot must be an administrator in the supergroup or channel for this to work and must have the appropriate administrator rights
// Returns True on success.
type BanChatSenderChat struct {
	ChatID       any   `json:"chat_id"`
	SenderChatID int64 `json:"sender_chat_id"`
}

func (entity *BanChatSenderChat) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *BanChatSenderChat) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity BanChatSenderChat) MarshalJSON() ([]byte, error) {
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
	type x0 BanChatSenderChat
	return json.Marshal((x0)(entity))
}

func (BanChatSenderChat) MethodName() string {
	return "banChatSenderChat"
}

func (BanChatSenderChat) ParseResult(response []byte) (*rawTypes.Result, error) {
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
