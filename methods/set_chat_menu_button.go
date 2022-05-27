package methods

import "github.com/Squirrel-Network/gobotapi/types"
import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "encoding/json"

type SetChatMenuButton struct {
	ChatId int64 `json:"chat_id,omitempty"`
	MenuButton *types.MenuButton `json:"menu_button,omitempty"`
}

func (entity *SetChatMenuButton) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (SetChatMenuButton) MethodName() string {
	return "setChatMenuButton"
}

func (SetChatMenuButton) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result bool `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeBoolean,
		Result: x1.Result,
	}
	return &result, nil
}
