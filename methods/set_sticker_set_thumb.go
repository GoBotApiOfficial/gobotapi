package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

// SetStickerSetThumb Use this method to set the thumbnail of a sticker set
// Animated thumbnails can be set for animated sticker sets only
// Video thumbnails can be set only for video sticker sets only
// Returns True on success.
type SetStickerSetThumb struct {
	Name string `json:"name"`
	Thumb rawTypes.InputFile `json:"thumb,omitempty"`
	UserId int64 `json:"user_id"`
}

func (entity *SetStickerSetThumb) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Thumb.(type) {
		case types.InputFile:
			files["thumb"] = entity.Thumb
			entity.Thumb = types.InputURL("attach://thumb")
	}
	return files
}

func (SetStickerSetThumb) MethodName() string {
	return "setStickerSetThumb"
}

func (SetStickerSetThumb) ParseResult(response []byte) (*rawTypes.Result, error) {
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
