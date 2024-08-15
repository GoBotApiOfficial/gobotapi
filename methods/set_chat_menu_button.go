// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"reflect"
)

// SetChatMenuButton Use this method to change the bot's menu button in a private chat, or the default menu button
// Returns True on success.
type SetChatMenuButton struct {
	ChatID     int64             `json:"chat_id,omitempty"`
	MenuButton *types.MenuButton `json:"menu_button,omitempty"`
}

func (entity *SetChatMenuButton) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SetChatMenuButton) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity SetChatMenuButton) MarshalJSON() ([]byte, error) {
	if reflect.DeepEqual(entity.MenuButton, nil) {
		entity.MenuButton = nil
	}
	type x0 SetChatMenuButton
	return json.Marshal((x0)(entity))
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
	result := rawTypes.Result{
		Kind:   types.TypeBoolean,
		Result: x1.Result,
	}
	return &result, nil
}
