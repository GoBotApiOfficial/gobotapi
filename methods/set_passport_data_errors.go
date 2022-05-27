package methods

import "github.com/Squirrel-Network/gobotapi/types"
import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "encoding/json"

type SetPassportDataErrors struct {
	Errors []types.PassportElementError `json:"errors,omitempty"`
	UserId int64 `json:"user_id"`
}

func (entity *SetPassportDataErrors) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (SetPassportDataErrors) MethodName() string {
	return "setPassportDataErrors"
}

func (SetPassportDataErrors) ParseResult(response []byte) (*rawTypes.Result, error) {
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
