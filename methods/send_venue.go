// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"reflect"
)

// SendVenue Use this method to send information about a venue
// On success, the sent Message is returned.
type SendVenue struct {
	Address              string                 `json:"address"`
	AllowPaidBroadcast   bool                   `json:"allow_paid_broadcast,omitempty"`
	BusinessConnectionID string                 `json:"business_connection_id,omitempty"`
	ChatID               any                    `json:"chat_id"`
	DisableNotification  bool                   `json:"disable_notification,omitempty"`
	FoursquareID         string                 `json:"foursquare_id,omitempty"`
	FoursquareType       string                 `json:"foursquare_type,omitempty"`
	GooglePlaceID        string                 `json:"google_place_id,omitempty"`
	GooglePlaceType      string                 `json:"google_place_type,omitempty"`
	Latitude             float64                `json:"latitude"`
	Longitude            float64                `json:"longitude"`
	MessageEffectID      string                 `json:"message_effect_id,omitempty"`
	MessageThreadID      int64                  `json:"message_thread_id,omitempty"`
	ProtectContent       bool                   `json:"protect_content,omitempty"`
	ReplyMarkup          any                    `json:"reply_markup,omitempty"`
	ReplyParameters      *types.ReplyParameters `json:"reply_parameters,omitempty"`
	Title                string                 `json:"title"`
}

func (entity *SendVenue) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SendVenue) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity SendVenue) MarshalJSON() ([]byte, error) {
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
	if entity.ReplyMarkup != nil {
		switch entity.ReplyMarkup.(type) {
		case *types.InlineKeyboardMarkup, *types.ReplyKeyboardMarkup, *types.ReplyKeyboardRemove, *types.ForceReply:
			break
		default:
			return nil, fmt.Errorf("reply_markup: unknown type: %T", entity.ReplyMarkup)
		}
	}
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 SendVenue
	return json.Marshal((x0)(entity))
}

func (SendVenue) MethodName() string {
	return "sendVenue"
}

func (SendVenue) ParseResult(response []byte) (*rawTypes.Result, error) {
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
