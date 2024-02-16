// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"gobotapi/types"
	rawTypes "gobotapi/types/raw"
)

// GetFile Use this method to get basic information about a file and prepare it for downloading
// For the moment, bots can download files of up to 20MB in size
// On success, a File object is returned
// The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response
// It is guaranteed that the link will be valid for at least 1 hour
// When the link expires, a new one can be requested by calling getFile again.
// Note: This function may not preserve the original file name and MIME type
// You should save the file's MIME type and name (if available) when the File object is received.
type GetFile struct {
	FileID string `json:"file_id"`
}

func (entity *GetFile) ProgressCallable() rawTypes.ProgressCallable {
	return nil
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
	result := rawTypes.Result{
		Kind:   types.TypeFile,
		Result: x1.Result,
	}
	return &result, nil
}
