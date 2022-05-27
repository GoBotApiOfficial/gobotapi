package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

type GetFile struct {
	FileId string `json:"file_id"`
}

func (entity *GetFile) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (GetFile) MethodName() string {
	return "getFile"
}

func (GetFile) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.File `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeFile,
		Result: x1.Result,
	}
	return &result, nil
}
