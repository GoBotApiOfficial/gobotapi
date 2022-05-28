package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

// LogOut Use this method to log out from the cloud Bot API server before launching the bot locally
// You must log out the bot before running it locally, otherwise there is no guarantee that the bot will receive updates
// After a successful call, you can immediately log in on a local server, but will not be able to log in back to the cloud Bot API server for 10 minutes
// Returns True on success
// Requires no parameters.
type LogOut struct{}

func (entity *LogOut) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (LogOut) MethodName() string {
	return "logOut"
}

func (LogOut) ParseResult(response []byte) (*rawTypes.Result, error) {
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
