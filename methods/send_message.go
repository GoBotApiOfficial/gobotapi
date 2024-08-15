// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"reflect"
)

// SendMessage Use this method to send text messages
// On success, the sent Message is returned.
type SendMessage struct {
	BusinessConnectionID string                    `json:"business_connection_id,omitempty"`
	ChatID               any                       `json:"chat_id"`
	DisableNotification  bool                      `json:"disable_notification,omitempty"`
	Entities             []types.MessageEntity     `json:"entities,omitempty"`
	LinkPreviewOptions   *types.LinkPreviewOptions `json:"link_preview_options,omitempty"`
	MessageEffectID      string                    `json:"message_effect_id,omitempty"`
	MessageThreadID      int64                     `json:"message_thread_id,omitempty"`
	ParseMode            string                    `json:"parse_mode,omitempty"`
	ProtectContent       bool                      `json:"protect_content,omitempty"`
	ReplyMarkup          any                       `json:"reply_markup,omitempty"`
	ReplyParameters      *types.ReplyParameters    `json:"reply_parameters,omitempty"`
	Text                 string                    `json:"text"`
}

func (entity *SendMessage) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SendMessage) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity SendMessage) MarshalJSON() ([]byte, error) {
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
	if nilCheck(entity.LinkPreviewOptions) {
		entity.LinkPreviewOptions = nil
	}
	if nilCheck(entity.ReplyParameters) {
		entity.ReplyParameters = nil
	}
	if nilCheck(entity.ReplyMarkup) {
		entity.ReplyMarkup = nil
	}
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	if entity.ReplyMarkup != nil {
		switch entity.ReplyMarkup.(type) {
		case *types.InlineKeyboardMarkup, *types.ReplyKeyboardMarkup, *types.ReplyKeyboardRemove, *types.ForceReply:
			break
		default:
			return nil, fmt.Errorf("reply_markup: unknown type: %T", entity.ReplyMarkup)
		}
	}
	type x0 SendMessage
	return json.Marshal((x0)(entity))
}

func (SendMessage) MethodName() string {
	return "sendMessage"
}

func (SendMessage) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.Message `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeMessage,
		Result: x1.Result,
	}
	return &result, nil
}
