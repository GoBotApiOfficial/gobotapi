// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"gobotapi/types"
	rawTypes "gobotapi/types/raw"
)

// SendMediaGroup Use this method to send a group of photos, videos, documents or audios as an album
// Documents and audio files can be only grouped in an album with messages of the same type
// On success, an array of Messages that were sent is returned.
type SendMediaGroup struct {
	BusinessConnectionID string                    `json:"business_connection_id,omitempty"`
	ChatID               any                       `json:"chat_id"`
	DisableNotification  bool                      `json:"disable_notification,omitempty"`
	Media                []types.InputMedia        `json:"media,omitempty"`
	MessageThreadID      int64                     `json:"message_thread_id,omitempty"`
	ProtectContent       bool                      `json:"protect_content,omitempty"`
	ReplyParameters      *types.ReplyParameters    `json:"reply_parameters,omitempty"`
	Progress             rawTypes.ProgressCallable `json:"-"`
}

func (entity *SendMediaGroup) ProgressCallable() rawTypes.ProgressCallable {
	return entity.Progress
}

func (entity *SendMediaGroup) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	for i, x0 := range entity.Media {
		x1 := x0.(rawTypes.InputMediaFiles).Files()
		for k, v := range x1 {
			var attachName string
			if k == "thumbnail" {
				attachName = fmt.Sprintf("file-%d-thumbnail", i)
				x0.SetAttachmentThumb(attachName)
			} else {
				attachName = fmt.Sprintf("file-%d", i)
				x0.SetAttachment(attachName)
			}
			files[attachName] = v
		}
	}
	return files
}

func (entity SendMediaGroup) MarshalJSON() ([]byte, error) {
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	for _, x0 := range entity.Media {
		switch x0.(type) {
		case *types.InputMediaAudio, *types.InputMediaDocument, *types.InputMediaPhoto, *types.InputMediaVideo:
			break
		default:
			return nil, fmt.Errorf("media: unknown type: %T", x0)
		}
	}
	type x0 SendMediaGroup
	return json.Marshal((x0)(entity))
}

func (SendMediaGroup) MethodName() string {
	return "sendMediaGroup"
}

func (SendMediaGroup) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result []types.Message `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeArrayOfMessage,
		Result: x1.Result,
	}
	return &result, nil
}
