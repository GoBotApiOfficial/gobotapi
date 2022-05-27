package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"
import "fmt"

type SendVideoNote struct {
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	ChatId int64 `json:"chat_id"`
	DisableNotification bool `json:"disable_notification,omitempty"`
	Duration int `json:"duration,omitempty"`
	Length int `json:"length,omitempty"`
	ProtectContent bool `json:"protect_content,omitempty"`
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
	ReplyToMessageId int64 `json:"reply_to_message_id,omitempty"`
	Thumb rawTypes.InputFile `json:"thumb,omitempty"`
	VideoNote rawTypes.InputFile `json:"video_note,omitempty"`
}

func (entity *SendVideoNote) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Thumb.(type) {
		case types.InputFile:
			files["thumb"] = entity.Thumb
			entity.Thumb = types.InputURL("attach://thumb")
	}
	switch entity.VideoNote.(type) {
		case types.InputFile:
			files["video_note"] = entity.VideoNote
			entity.VideoNote = nil
	}
	return files
}

func (entity SendVideoNote) MarshalJSON() ([]byte, error) {
	if entity.ReplyMarkup != nil {
		switch entity.ReplyMarkup.(type) {
			case *types.InlineKeyboardMarkup, *types.ReplyKeyboardMarkup, *types.ReplyKeyboardRemove, *types.ForceReply:
				break
			default:
				return nil, fmt.Errorf("reply_markup: unknown type: %T", entity.ReplyMarkup)
		}
	}
	type x0 SendVideoNote
	return json.Marshal((x0)(entity))
}

func (SendVideoNote) MethodName() string {
	return "sendVideoNote"
}

func (SendVideoNote) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.Message `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeMessage,
		Result: x1.Result,
	}
	return &result, nil
}
