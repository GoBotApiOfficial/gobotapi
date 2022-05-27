package methods

import "github.com/Squirrel-Network/gobotapi/types"
import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "encoding/json"

type StopPoll struct {
	ChatId int64 `json:"chat_id"`
	MessageId int64 `json:"message_id"`
	ReplyMarkup *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (entity *StopPoll) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (StopPoll) MethodName() string {
	return "stopPoll"
}

func (StopPoll) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.Poll `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypePoll,
		Result: x1.Result,
	}
	return &result, nil
}
