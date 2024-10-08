// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"reflect"
)

// ForwardMessages Use this method to forward multiple messages of any kind
// If some of the specified messages can't be found or forwarded, they are skipped
// Service messages and messages with protected content can't be forwarded
// Album grouping is kept for forwarded messages
// On success, an array of MessageId of the sent messages is returned.
type ForwardMessages struct {
	ChatID              any     `json:"chat_id"`
	DisableNotification bool    `json:"disable_notification,omitempty"`
	FromChatID          int64   `json:"from_chat_id"`
	MessageIDs          []int64 `json:"message_ids,omitempty"`
	MessageThreadID     int64   `json:"message_thread_id,omitempty"`
	ProtectContent      bool    `json:"protect_content,omitempty"`
}

func (entity *ForwardMessages) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *ForwardMessages) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity ForwardMessages) MarshalJSON() ([]byte, error) {
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
	type x0 ForwardMessages
	return json.Marshal((x0)(entity))
}

func (ForwardMessages) MethodName() string {
	return "forwardMessages"
}

func (ForwardMessages) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result []types.MessageId `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeArrayOfMessageId,
		Result: x1.Result,
	}
	return &result, nil
}
