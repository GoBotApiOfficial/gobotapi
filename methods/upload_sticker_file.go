// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// UploadStickerFile Use this method to upload a file with a sticker for later use in the createNewStickerSet and addStickerToSet methods (the file can be used multiple times)
// Returns the uploaded File on success.
type UploadStickerFile struct {
	Sticker       rawTypes.InputFile        `json:"sticker,omitempty"`
	StickerFormat string                    `json:"sticker_format"`
	UserID        int64                     `json:"user_id"`
	Progress      rawTypes.ProgressCallable `json:"-"`
}

func (entity *UploadStickerFile) ProgressCallable() rawTypes.ProgressCallable {
	return entity.Progress
}

func (entity *UploadStickerFile) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Sticker.(type) {
	case types.InputBytes:
		files["sticker"] = entity.Sticker
		entity.Sticker = nil
	}
	return files
}

func (UploadStickerFile) MethodName() string {
	return "uploadStickerFile"
}

func (UploadStickerFile) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.File `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeFile,
		Result: x1.Result,
	}
	return &result, nil
}
