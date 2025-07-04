// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"reflect"
)

// SendChecklist Use this method to send a checklist on behalf of a connected business account
// On success, the sent Message is returned.
type SendChecklist struct {
	BusinessConnectionID string                      `json:"business_connection_id"`
	ChatID               int64                       `json:"chat_id"`
	Checklist            types.InputChecklist        `json:"checklist"`
	DisableNotification  bool                        `json:"disable_notification,omitempty"`
	MessageEffectID      string                      `json:"message_effect_id,omitempty"`
	ProtectContent       bool                        `json:"protect_content,omitempty"`
	ReplyMarkup          *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ReplyParameters      *types.ReplyParameters      `json:"reply_parameters,omitempty"`
}

func (entity *SendChecklist) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SendChecklist) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity SendChecklist) MarshalJSON() ([]byte, error) {
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
	type x0 SendChecklist
	return json.Marshal((x0)(entity))
}

func (SendChecklist) MethodName() string {
	return "sendChecklist"
}

func (SendChecklist) ParseResult(response []byte) (*rawTypes.Result, error) {
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
