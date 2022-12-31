// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

// SendPoll Use this method to send a native poll
// On success, the sent Message is returned.
type SendPoll struct {
	AllowSendingWithoutReply bool                  `json:"allow_sending_without_reply,omitempty"`
	AllowsMultipleAnswers    bool                  `json:"allows_multiple_answers,omitempty"`
	ChatID                   any                   `json:"chat_id"`
	CloseDate                int64                 `json:"close_date,omitempty"`
	CorrectOptionID          int64                 `json:"correct_option_id,omitempty"`
	DisableNotification      bool                  `json:"disable_notification,omitempty"`
	Explanation              string                `json:"explanation,omitempty"`
	ExplanationEntities      []types.MessageEntity `json:"explanation_entities,omitempty"`
	ExplanationParseMode     string                `json:"explanation_parse_mode,omitempty"`
	IsAnonymous              bool                  `json:"is_anonymous,omitempty"`
	IsClosed                 bool                  `json:"is_closed,omitempty"`
	MessageThreadID          int64                 `json:"message_thread_id,omitempty"`
	OpenPeriod               int                   `json:"open_period,omitempty"`
	Options                  []string              `json:"options,omitempty"`
	ProtectContent           bool                  `json:"protect_content,omitempty"`
	Question                 string                `json:"question"`
	ReplyMarkup              any                   `json:"reply_markup,omitempty"`
	ReplyToMessageID         int64                 `json:"reply_to_message_id,omitempty"`
	Type                     string                `json:"type,omitempty"`
}

func (entity *SendPoll) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SendPoll) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity SendPoll) MarshalJSON() ([]byte, error) {
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
	type x0 SendPoll
	return json.Marshal((x0)(entity))
}

func (SendPoll) MethodName() string {
	return "sendPoll"
}

func (SendPoll) ParseResult(response []byte) (*rawTypes.Result, error) {
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
