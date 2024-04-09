// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"gobotapi/types"
	rawTypes "gobotapi/types/raw"
)

// ReplaceStickerInSet Use this method to replace an existing sticker in a sticker set with a new one
// The method is equivalent to calling deleteStickerFromSet, then addStickerToSet, then setStickerPositionInSet
// Returns True on success.
type ReplaceStickerInSet struct {
	Name       string             `json:"name"`
	OldSticker string             `json:"old_sticker"`
	Sticker    types.InputSticker `json:"sticker"`
	UserID     int64              `json:"user_id"`
}

func (entity *ReplaceStickerInSet) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *ReplaceStickerInSet) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (ReplaceStickerInSet) MethodName() string {
	return "replaceStickerInSet"
}

func (ReplaceStickerInSet) ParseResult(response []byte) (*rawTypes.Result, error) {
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
