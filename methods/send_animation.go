package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"
import "fmt"

// SendAnimation Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound)
// On success, the sent Message is returned
// Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
type SendAnimation struct {
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	Animation rawTypes.InputFile `json:"animation,omitempty"`
	Caption string `json:"caption,omitempty"`
	CaptionEntities []types.MessageEntity `json:"caption_entities,omitempty"`
	ChatId int64 `json:"chat_id"`
	DisableNotification bool `json:"disable_notification,omitempty"`
	Duration int `json:"duration,omitempty"`
	Height int `json:"height,omitempty"`
	ParseMode string `json:"parse_mode,omitempty"`
	ProtectContent bool `json:"protect_content,omitempty"`
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
	ReplyToMessageId int64 `json:"reply_to_message_id,omitempty"`
	Thumb rawTypes.InputFile `json:"thumb,omitempty"`
	Width int64 `json:"width,omitempty"`
}

func (entity *SendAnimation) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Animation.(type) {
		case types.InputFile:
			files["animation"] = entity.Animation
			entity.Animation = nil
	}
	switch entity.Thumb.(type) {
		case types.InputFile:
			files["thumb"] = entity.Thumb
			entity.Thumb = types.InputURL("attach://thumb")
	}
	return files
}

func (entity SendAnimation) MarshalJSON() ([]byte, error) {
	if entity.ReplyMarkup != nil {
		switch entity.ReplyMarkup.(type) {
			case *types.InlineKeyboardMarkup, *types.ReplyKeyboardMarkup, *types.ReplyKeyboardRemove, *types.ForceReply:
				break
			default:
				return nil, fmt.Errorf("reply_markup: unknown type: %T", entity.ReplyMarkup)
		}
	}
	type x0 SendAnimation
	return json.Marshal((x0)(entity))
}

func (SendAnimation) MethodName() string {
	return "sendAnimation"
}

func (SendAnimation) ParseResult(response []byte) (*rawTypes.Result, error) {
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
