// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// SendGift Sends a gift to the given user
// The gift can't be converted to Telegram Stars by the user
// Returns True on success.
type SendGift struct {
	GiftID        string                `json:"gift_id"`
	PayForUpgrade bool                  `json:"pay_for_upgrade,omitempty"`
	Text          string                `json:"text,omitempty"`
	TextEntities  []types.MessageEntity `json:"text_entities,omitempty"`
	TextParseMode string                `json:"text_parse_mode,omitempty"`
	UserID        int64                 `json:"user_id"`
}

func (entity *SendGift) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SendGift) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (SendGift) MethodName() string {
	return "sendGift"
}

func (SendGift) ParseResult(response []byte) (*rawTypes.Result, error) {
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
