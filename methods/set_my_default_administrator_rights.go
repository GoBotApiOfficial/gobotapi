package methods

import "github.com/Squirrel-Network/gobotapi/types"
import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "encoding/json"

// SetMyDefaultAdministratorRights Use this method to change the default administrator rights requested by the bot when it's added as an administrator to groups or channels
// These rights will be suggested to users, but they are are free to modify the list before adding the bot
// Returns True on success.
type SetMyDefaultAdministratorRights struct {
	ForChannels bool `json:"for_channels,omitempty"`
	Rights *types.ChatAdministratorRights `json:"rights,omitempty"`
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
	result := rawTypes.Result {
		Kind: types.TypeBoolean,
		Result: x1.Result,
	}
	return &result, nil
}
