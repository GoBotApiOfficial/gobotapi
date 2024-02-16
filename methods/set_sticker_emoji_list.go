// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"gobotapi/types"
	rawTypes "gobotapi/types/raw"
)

// SetStickerEmojiList Use this method to change the list of emoji assigned to a regular or custom emoji sticker
// The sticker must belong to a sticker set created by the bot
// Returns True on success.
type SetStickerEmojiList struct {
	EmojiList []string `json:"emoji_list,omitempty"`
	Sticker   string   `json:"sticker"`
}

func (entity *SetStickerEmojiList) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SetStickerEmojiList) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (SetStickerEmojiList) MethodName() string {
	return "setStickerEmojiList"
}

func (SetStickerEmojiList) ParseResult(response []byte) (*rawTypes.Result, error) {
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
