package methods

import "github.com/Squirrel-Network/gobotapi/types"
import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "encoding/json"
import "fmt"

type SendVideo struct {
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
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
	SupportsStreaming bool `json:"supports_streaming,omitempty"`
	Thumb rawTypes.InputFile `json:"thumb,omitempty"`
	Video rawTypes.InputFile `json:"video,omitempty"`
	Width int64 `json:"width,omitempty"`
}

func (entity *SendVideo) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Thumb.(type) {
		case types.InputFile:
			files["thumb"] = entity.Thumb
			entity.Thumb = types.InputURL("attach://thumb")
	}
	switch entity.Video.(type) {
		case types.InputFile:
			files["video"] = entity.Video
			entity.Video = nil
	}
	return files
}

func (entity SendVideo) MarshalJSON() ([]byte, error) {
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
	result := rawTypes.Result {
		Kind: types.TypeMessage,
		Result: x1.Result,
	}
	return &result, nil
}
