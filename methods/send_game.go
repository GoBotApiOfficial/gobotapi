// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"reflect"
)

// SendGame Use this method to send a game
// On success, the sent Message is returned.
type SendGame struct {
	BusinessConnectionID string                      `json:"business_connection_id,omitempty"`
	ChatID               int64                       `json:"chat_id"`
	DisableNotification  bool                        `json:"disable_notification,omitempty"`
	GameShortName        string                      `json:"game_short_name"`
	MessageEffectID      string                      `json:"message_effect_id,omitempty"`
	MessageThreadID      int64                       `json:"message_thread_id,omitempty"`
	ProtectContent       bool                        `json:"protect_content,omitempty"`
	ReplyMarkup          *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ReplyParameters      *types.ReplyParameters      `json:"reply_parameters,omitempty"`
}

func (entity *SendGame) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SendGame) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity SendGame) MarshalJSON() ([]byte, error) {
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
	if nilCheck(entity.ReplyParameters) {
		entity.ReplyParameters = nil
	}
	if nilCheck(entity.ReplyMarkup) {
		entity.ReplyMarkup = nil
	}
	type x0 SendGame
	return json.Marshal((x0)(entity))
}

func (SendGame) MethodName() string {
	return "sendGame"
}

func (SendGame) ParseResult(response []byte) (*rawTypes.Result, error) {
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
