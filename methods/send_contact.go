package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "encoding/json"
import "fmt"
import "github.com/Squirrel-Network/gobotapi/types"

// SendContact Use this method to send phone contacts
// On success, the sent Message is returned.
type SendContact struct {
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	ChatId int64 `json:"chat_id"`
	DisableNotification bool `json:"disable_notification,omitempty"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number"`
	ProtectContent bool `json:"protect_content,omitempty"`
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
	ReplyToMessageId int64 `json:"reply_to_message_id,omitempty"`
	Vcard string `json:"vcard,omitempty"`
}

func (entity *SendContact) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity SendContact) MarshalJSON() ([]byte, error) {
	if entity.ReplyMarkup != nil {
		switch entity.ReplyMarkup.(type) {
			case *types.InlineKeyboardMarkup, *types.ReplyKeyboardMarkup, *types.ReplyKeyboardRemove, *types.ForceReply:
				break
			default:
				return nil, fmt.Errorf("reply_markup: unknown type: %T", entity.ReplyMarkup)
		}
	}
	type x0 SendContact
	return json.Marshal((x0)(entity))
}

func (SendContact) MethodName() string {
	return "sendContact"
}

func (SendContact) ParseResult(response []byte) (*rawTypes.Result, error) {
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
