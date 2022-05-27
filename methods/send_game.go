package methods

import "github.com/Squirrel-Network/gobotapi/types"
import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "encoding/json"

type SendGame struct {
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	ChatId int64 `json:"chat_id"`
	DisableNotification bool `json:"disable_notification,omitempty"`
	GameShortName string `json:"game_short_name"`
	ProtectContent bool `json:"protect_content,omitempty"`
	ReplyMarkup *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ReplyToMessageId int64 `json:"reply_to_message_id,omitempty"`
}

func (entity *SendGame) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (SendGame) MethodName() string {
	return "sendGame"
}

func (SendGame) ParseResult(response []byte) (*rawTypes.Result, error) {
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
