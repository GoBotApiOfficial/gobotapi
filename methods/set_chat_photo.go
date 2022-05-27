package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

type SetChatPhoto struct {
	ChatId int64 `json:"chat_id"`
	Photo rawTypes.InputFile `json:"photo,omitempty"`
}

func (entity *SetChatPhoto) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Photo.(type) {
		case types.InputFile:
			files["photo"] = entity.Photo
			entity.Photo = nil
	}
	return files
}

func (SetChatPhoto) MethodName() string {
	return "setChatPhoto"
}

func (SetChatPhoto) ParseResult(response []byte) (*rawTypes.Result, error) {
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
