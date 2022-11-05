// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

// SendPhoto Use this method to send photos
// On success, the sent Message is returned.
type SendPhoto struct {
	AllowSendingWithoutReply bool                      `json:"allow_sending_without_reply,omitempty"`
	Caption                  string                    `json:"caption,omitempty"`
	CaptionEntities          []types.MessageEntity     `json:"caption_entities,omitempty"`
	ChatID                   any                       `json:"chat_id"`
	DisableNotification      bool                      `json:"disable_notification,omitempty"`
	MessageThreadID          int64                     `json:"message_thread_id,omitempty"`
	ParseMode                string                    `json:"parse_mode,omitempty"`
	Photo                    rawTypes.InputFile        `json:"photo,omitempty"`
	ProtectContent           bool                      `json:"protect_content,omitempty"`
	ReplyMarkup              any                       `json:"reply_markup,omitempty"`
	ReplyToMessageID         int64                     `json:"reply_to_message_id,omitempty"`
	Progress                 rawTypes.ProgressCallable `json:"-"`
}

func (entity *SendPhoto) ProgressCallable() rawTypes.ProgressCallable {
	return entity.Progress
}

func (entity *SendPhoto) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Photo.(type) {
	case types.InputBytes:
		files["photo"] = entity.Photo
		entity.Photo = nil
	}
	return files
}

func (entity SendPhoto) MarshalJSON() ([]byte, error) {
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
	type x0 SendPhoto
	return json.Marshal((x0)(entity))
}

func (SendPhoto) MethodName() string {
	return "sendPhoto"
}

func (SendPhoto) ParseResult(response []byte) (*rawTypes.Result, error) {
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
