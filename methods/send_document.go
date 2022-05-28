package methods

import "github.com/Squirrel-Network/gobotapi/types"
import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "encoding/json"
import "fmt"

// SendDocument Use this method to send general files
// On success, the sent Message is returned
// Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
type SendDocument struct {
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	Caption string `json:"caption,omitempty"`
	CaptionEntities []types.MessageEntity `json:"caption_entities,omitempty"`
	ChatId int64 `json:"chat_id"`
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
	DisableNotification bool `json:"disable_notification,omitempty"`
	Document rawTypes.InputFile `json:"document,omitempty"`
	ParseMode string `json:"parse_mode,omitempty"`
	ProtectContent bool `json:"protect_content,omitempty"`
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
	ReplyToMessageId int64 `json:"reply_to_message_id,omitempty"`
	Thumb rawTypes.InputFile `json:"thumb,omitempty"`
}

func (entity *SendDocument) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Document.(type) {
		case types.InputFile:
			files["document"] = entity.Document
			entity.Document = nil
	}
	switch entity.Thumb.(type) {
		case types.InputFile:
			files["thumb"] = entity.Thumb
			entity.Thumb = types.InputURL("attach://thumb")
	}
	return files
}

func (entity SendDocument) MarshalJSON() ([]byte, error) {
	if entity.ReplyMarkup != nil {
		switch entity.ReplyMarkup.(type) {
			case *types.InlineKeyboardMarkup, *types.ReplyKeyboardMarkup, *types.ReplyKeyboardRemove, *types.ForceReply:
				break
			default:
				return nil, fmt.Errorf("reply_markup: unknown type: %T", entity.ReplyMarkup)
		}
	}
	type x0 SendDocument
	return json.Marshal((x0)(entity))
}

func (SendDocument) MethodName() string {
	return "sendDocument"
}

func (SendDocument) ParseResult(response []byte) (*rawTypes.Result, error) {
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
