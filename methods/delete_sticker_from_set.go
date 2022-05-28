package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

// DeleteStickerFromSet Use this method to delete a sticker from a set created by the bot
// Returns True on success.
type DeleteStickerFromSet struct {
	Sticker string `json:"sticker"`
}

func (entity *DeleteStickerFromSet) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (DeleteStickerFromSet) MethodName() string {
	return "deleteStickerFromSet"
}

func (DeleteStickerFromSet) ParseResult(response []byte) (*rawTypes.Result, error) {
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
