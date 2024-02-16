// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"gobotapi/types"
	rawTypes "gobotapi/types/raw"
)

// SetStickerMaskPosition Use this method to change the mask position of a mask sticker
// The sticker must belong to a sticker set that was created by the bot
// Returns True on success.
type SetStickerMaskPosition struct {
	MaskPosition *types.MaskPosition `json:"mask_position,omitempty"`
	Sticker      string              `json:"sticker"`
}

func (entity *SetStickerMaskPosition) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SetStickerMaskPosition) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (SetStickerMaskPosition) MethodName() string {
	return "setStickerMaskPosition"
}

func (SetStickerMaskPosition) ParseResult(response []byte) (*rawTypes.Result, error) {
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
