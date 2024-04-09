// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// SetStickerSetTitle Use this method to set the title of a created sticker set
// Returns True on success.
type SetStickerSetTitle struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

func (entity *SetStickerSetTitle) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SetStickerSetTitle) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (SetStickerSetTitle) MethodName() string {
	return "setStickerSetTitle"
}

func (SetStickerSetTitle) ParseResult(response []byte) (*rawTypes.Result, error) {
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
