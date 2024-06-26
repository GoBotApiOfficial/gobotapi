// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// SetMyDefaultAdministratorRights Use this method to change the default administrator rights requested by the bot when it's added as an administrator to groups or channels
// These rights will be suggested to users, but they are free to modify the list before adding the bot
// Returns True on success.
type SetMyDefaultAdministratorRights struct {
	ForChannels bool                           `json:"for_channels,omitempty"`
	Rights      *types.ChatAdministratorRights `json:"rights,omitempty"`
}

func (entity *SetMyDefaultAdministratorRights) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SetMyDefaultAdministratorRights) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (SetMyDefaultAdministratorRights) MethodName() string {
	return "setMyDefaultAdministratorRights"
}

func (SetMyDefaultAdministratorRights) ParseResult(response []byte) (*rawTypes.Result, error) {
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
