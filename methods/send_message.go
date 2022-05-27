package methods

import "github.com/Squirrel-Network/gobotapi/types"
import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "encoding/json"
import "fmt"

type SendMessage struct {
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	ChatId int64 `json:"chat_id"`
	DisableNotification bool `json:"disable_notification,omitempty"`
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
	Entities []types.MessageEntity `json:"entities,omitempty"`
	ParseMode string `json:"parse_mode,omitempty"`
	ProtectContent bool `json:"protect_content,omitempty"`
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
	ReplyToMessageId int64 `json:"reply_to_message_id,omitempty"`
	Text string `json:"text"`
}

func (entity *SendMessage) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity SendMessage) MarshalJSON() ([]byte, error) {
	if entity.ReplyMarkup != nil {
		switch entity.ReplyMarkup.(type) {
			case *types.InlineKeyboardMarkup, *types.ReplyKeyboardMarkup, *types.ReplyKeyboardRemove, *types.ForceReply:
				break
			default:
				return nil, fmt.Errorf("reply_markup: unknown type: %T", entity.ReplyMarkup)
		}
	}
	type x0 SendMessage
	return json.Marshal((x0)(entity))
}

func (SendMessage) MethodName() string {
	return "sendMessage"
}

func (SendMessage) ParseResult(response []byte) (*rawTypes.Result, error) {
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
