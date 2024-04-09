// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// SendVideo Use this method to send video files, Telegram clients support MPEG4 videos (other formats may be sent as Document)
// On success, the sent Message is returned
// Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
type SendVideo struct {
	BusinessConnectionID string                    `json:"business_connection_id,omitempty"`
	Caption              string                    `json:"caption,omitempty"`
	CaptionEntities      []types.MessageEntity     `json:"caption_entities,omitempty"`
	ChatID               any                       `json:"chat_id"`
	DisableNotification  bool                      `json:"disable_notification,omitempty"`
	Duration             int                       `json:"duration,omitempty"`
	HasSpoiler           bool                      `json:"has_spoiler,omitempty"`
	Height               int                       `json:"height,omitempty"`
	MessageThreadID      int64                     `json:"message_thread_id,omitempty"`
	ParseMode            string                    `json:"parse_mode,omitempty"`
	ProtectContent       bool                      `json:"protect_content,omitempty"`
	ReplyMarkup          any                       `json:"reply_markup,omitempty"`
	ReplyParameters      *types.ReplyParameters    `json:"reply_parameters,omitempty"`
	SupportsStreaming    bool                      `json:"supports_streaming,omitempty"`
	Thumbnail            rawTypes.InputFile        `json:"thumbnail,omitempty"`
	Video                rawTypes.InputFile        `json:"video,omitempty"`
	Width                int64                     `json:"width,omitempty"`
	Progress             rawTypes.ProgressCallable `json:"-"`
}

func (entity *SendVideo) ProgressCallable() rawTypes.ProgressCallable {
	return entity.Progress
}

func (entity *SendVideo) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Thumbnail.(type) {
	case types.InputBytes:
		files["thumbnail"] = entity.Thumbnail
		entity.Thumbnail = types.InputURL("attach://thumbnail")
	}
	switch entity.Video.(type) {
	case types.InputBytes:
		files["video"] = entity.Video
		entity.Video = nil
	}
	return files
}

func (entity SendVideo) MarshalJSON() ([]byte, error) {
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
	type x0 SendVideo
	return json.Marshal((x0)(entity))
}

func (SendVideo) MethodName() string {
	return "sendVideo"
}

func (SendVideo) ParseResult(response []byte) (*rawTypes.Result, error) {
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
