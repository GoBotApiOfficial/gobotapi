package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

// GetChatMenuButton Use this method to get the current value of the bot's menu button in a private chat, or the default menu button
// Returns MenuButton on success.
type GetChatMenuButton struct {
	ChatId int64 `json:"chat_id,omitempty"`
}

func (entity *GetChatMenuButton) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (GetChatMenuButton) MethodName() string {
	return "getChatMenuButton"
}

func (GetChatMenuButton) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.MenuButton `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeMenuButton,
		Result: x1.Result,
	}
	return &result, nil
}
