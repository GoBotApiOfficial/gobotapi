package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

type SetGameScore struct {
	ChatId int64 `json:"chat_id,omitempty"`
	DisableEditMessage bool `json:"disable_edit_message,omitempty"`
	Force bool `json:"force,omitempty"`
	InlineMessageId string `json:"inline_message_id,omitempty"`
	MessageId int64 `json:"message_id,omitempty"`
	Score int `json:"score"`
	UserId int64 `json:"user_id"`
}

func (entity *SetGameScore) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (SetGameScore) MethodName() string {
	return "setGameScore"
}

func (SetGameScore) ParseResult(response []byte) (*rawTypes.Result, error) {
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
