// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"reflect"
)

// SendVoice Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message
// For this to work, your audio must be in an .OGG file encoded with OPUS, or in .MP3 format, or in .M4A format (other formats may be sent as Audio or Document)
// On success, the sent Message is returned
// Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
type SendVoice struct {
	AllowPaidBroadcast   bool                      `json:"allow_paid_broadcast,omitempty"`
	BusinessConnectionID string                    `json:"business_connection_id,omitempty"`
	Caption              string                    `json:"caption,omitempty"`
	CaptionEntities      []types.MessageEntity     `json:"caption_entities,omitempty"`
	ChatID               any                       `json:"chat_id"`
	DisableNotification  bool                      `json:"disable_notification,omitempty"`
	Duration             int                       `json:"duration,omitempty"`
	MessageEffectID      string                    `json:"message_effect_id,omitempty"`
	MessageThreadID      int64                     `json:"message_thread_id,omitempty"`
	ParseMode            string                    `json:"parse_mode,omitempty"`
	ProtectContent       bool                      `json:"protect_content,omitempty"`
	ReplyMarkup          any                       `json:"reply_markup,omitempty"`
	ReplyParameters      *types.ReplyParameters    `json:"reply_parameters,omitempty"`
	Voice                rawTypes.InputFile        `json:"voice,omitempty"`
	Progress             rawTypes.ProgressCallable `json:"-"`
}

func (entity *SendVoice) ProgressCallable() rawTypes.ProgressCallable {
	return entity.Progress
}

func (entity *SendVoice) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Voice.(type) {
	case types.InputBytes:
		files["voice"] = entity.Voice
		entity.Voice = nil
	}
	return files
}

func (entity SendVoice) MarshalJSON() ([]byte, error) {
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
	type x0 SendVoice
	return json.Marshal((x0)(entity))
}

func (SendVoice) MethodName() string {
	return "sendVoice"
}

func (SendVoice) ParseResult(response []byte) (*rawTypes.Result, error) {
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
