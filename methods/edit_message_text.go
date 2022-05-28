package methods

import "github.com/Squirrel-Network/gobotapi/types"
import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "encoding/json"

// EditMessageText Use this method to edit text and game messages
// On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageText struct {
	ChatId int64 `json:"chat_id,omitempty"`
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
	Entities []types.MessageEntity `json:"entities,omitempty"`
	InlineMessageId string `json:"inline_message_id,omitempty"`
	MessageId int64 `json:"message_id,omitempty"`
	ParseMode string `json:"parse_mode,omitempty"`
	ReplyMarkup *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Text string `json:"text"`
}

func (entity *EditMessageText) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (EditMessageText) MethodName() string {
	return "editMessageText"
}

func (EditMessageText) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x0 struct {
		Result bool `json:"result"`
	}
	_ = json.Unmarshal(response, &x0)
	if x0.Result {
		result := rawTypes.Result {
			Kind: types.TypeBoolean,
			Result: true,
		}
		return &result, nil
	} else {
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
}
