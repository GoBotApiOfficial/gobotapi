// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// SetCustomEmojiStickerSetThumbnail Use this method to set the thumbnail of a custom emoji sticker set
// Returns True on success.
type SetCustomEmojiStickerSetThumbnail struct {
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
	Name          string `json:"name"`
}

func (entity *SetCustomEmojiStickerSetThumbnail) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SetCustomEmojiStickerSetThumbnail) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (SetCustomEmojiStickerSetThumbnail) MethodName() string {
	return "setCustomEmojiStickerSetThumbnail"
}

func (SetCustomEmojiStickerSetThumbnail) ParseResult(response []byte) (*rawTypes.Result, error) {
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
