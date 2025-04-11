// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// SetBusinessAccountGiftSettings Changes the privacy settings pertaining to incoming gifts in a managed business account
// Requires the can_change_gift_settings business bot right
// Returns True on success.
type SetBusinessAccountGiftSettings struct {
	AcceptedGiftTypes    types.AcceptedGiftTypes `json:"accepted_gift_types"`
	BusinessConnectionID string                  `json:"business_connection_id"`
	ShowGiftButton       bool                    `json:"show_gift_button"`
}

func (entity *SetBusinessAccountGiftSettings) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SetBusinessAccountGiftSettings) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (SetBusinessAccountGiftSettings) MethodName() string {
	return "setBusinessAccountGiftSettings"
}

func (SetBusinessAccountGiftSettings) ParseResult(response []byte) (*rawTypes.Result, error) {
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
