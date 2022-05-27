package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

type ForwardMessage struct {
	ChatId int64 `json:"chat_id"`
	DisableNotification bool `json:"disable_notification,omitempty"`
	FromChatId int64 `json:"from_chat_id"`
	MessageId int64 `json:"message_id"`
	ProtectContent bool `json:"protect_content,omitempty"`
}

func (entity *ForwardMessage) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (ForwardMessage) MethodName() string {
	return "forwardMessage"
}

func (ForwardMessage) ParseResult(response []byte) (*rawTypes.Result, error) {
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
