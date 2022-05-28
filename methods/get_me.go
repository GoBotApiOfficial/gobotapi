package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

// GetMe A simple method for testing your bot's authentication token
// Requires no parameters
// Returns basic information about the bot in form of a User object.
type GetMe struct{}

func (entity *GetMe) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (GetMe) MethodName() string {
	return "getMe"
}

func (GetMe) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.User `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeUser,
		Result: x1.Result,
	}
	return &result, nil
}
