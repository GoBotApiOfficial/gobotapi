package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

// GetUpdates Use this method to receive incoming updates using long polling (wiki)
// An Array of Update objects is returned.
type GetUpdates struct {
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
	Limit int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
	Timeout int `json:"timeout,omitempty"`
}

func (entity *GetUpdates) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (GetUpdates) MethodName() string {
	return "getUpdates"
}

func (GetUpdates) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result []types.Update `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeArrayOfUpdate,
		Result: x1.Result,
	}
	return &result, nil
}
