package methods

import "github.com/Squirrel-Network/gobotapi/types"
import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "encoding/json"

type StopMessageLiveLocation struct {
	ChatId int64 `json:"chat_id,omitempty"`
	InlineMessageId string `json:"inline_message_id,omitempty"`
	MessageId int64 `json:"message_id,omitempty"`
	ReplyMarkup *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (entity *StopMessageLiveLocation) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (StopMessageLiveLocation) MethodName() string {
	return "stopMessageLiveLocation"
}

func (StopMessageLiveLocation) ParseResult(response []byte) (*rawTypes.Result, error) {
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
