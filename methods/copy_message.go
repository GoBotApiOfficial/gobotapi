package methods

import "github.com/Squirrel-Network/gobotapi/types"
import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "encoding/json"
import "fmt"

// CopyMessage Use this method to copy messages of any kind
// Service messages and invoice messages can't be copied
// The method is analogous to the method forwardMessage, but the copied message doesn't have a link to the original message
// Returns the MessageId of the sent message on success.
type CopyMessage struct {
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	Caption string `json:"caption,omitempty"`
	CaptionEntities []types.MessageEntity `json:"caption_entities,omitempty"`
	ChatId int64 `json:"chat_id"`
	DisableNotification bool `json:"disable_notification,omitempty"`
	FromChatId int64 `json:"from_chat_id"`
	MessageId int64 `json:"message_id"`
	ParseMode string `json:"parse_mode,omitempty"`
	ProtectContent bool `json:"protect_content,omitempty"`
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
	ReplyToMessageId int64 `json:"reply_to_message_id,omitempty"`
}

func (entity *CopyMessage) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity CopyMessage) MarshalJSON() ([]byte, error) {
	if entity.ReplyMarkup != nil {
		switch entity.ReplyMarkup.(type) {
			case *types.InlineKeyboardMarkup, *types.ReplyKeyboardMarkup, *types.ReplyKeyboardRemove, *types.ForceReply:
				break
			default:
				return nil, fmt.Errorf("reply_markup: unknown type: %T", entity.ReplyMarkup)
		}
	}
	type x0 CopyMessage
	return json.Marshal((x0)(entity))
}

func (CopyMessage) MethodName() string {
	return "copyMessage"
}

func (CopyMessage) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.MessageId `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeMessageId,
		Result: x1.Result,
	}
	return &result, nil
}
