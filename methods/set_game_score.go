// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// SetGameScore Use this method to set the score of the specified user in a game message
// On success, if the message is not an inline message, the Message is returned, otherwise True is returned
// Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
type SetGameScore struct {
	ChatID             int64  `json:"chat_id,omitempty"`
	DisableEditMessage bool   `json:"disable_edit_message,omitempty"`
	Force              bool   `json:"force,omitempty"`
	InlineMessageID    string `json:"inline_message_id,omitempty"`
	MessageID          int64  `json:"message_id,omitempty"`
	Score              int    `json:"score"`
	UserID             int64  `json:"user_id"`
}

func (entity *SetGameScore) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SetGameScore) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (SetGameScore) MethodName() string {
	return "setGameScore"
}

func (SetGameScore) ParseResult(response []byte) (*rawTypes.Result, error) {
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
