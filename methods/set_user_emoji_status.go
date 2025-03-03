// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// SetUserEmojiStatus Changes the emoji status for a given user that previously allowed the bot to manage their emoji status via the Mini App method requestEmojiStatusAccess
// Returns True on success.
type SetUserEmojiStatus struct {
	EmojiStatusCustomEmojiID  string `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate int64  `json:"emoji_status_expiration_date,omitempty"`
	UserID                    int64  `json:"user_id"`
}

func (entity *SetUserEmojiStatus) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SetUserEmojiStatus) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (SetUserEmojiStatus) MethodName() string {
	return "setUserEmojiStatus"
}

func (SetUserEmojiStatus) ParseResult(response []byte) (*rawTypes.Result, error) {
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
