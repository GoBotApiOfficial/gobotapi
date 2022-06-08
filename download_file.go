package gobotapi

import (
	"errors"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
	"os"
	"path"
)

func (ctx *Client) DownloadFile(fileId, filePath string, progress rawTypes.ProgressCallable) error {
	bytes, err := ctx.DownloadBytes(fileId, progress)
	if err != nil {
		return err
	}
	if res, err := os.Stat(path.Dir(filePath)); os.IsNotExist(err) {
		return errors.New("path not found")
	} else {
		_ = os.WriteFile(filePath, bytes, res.Mode().Perm())
	}
	return nil
}
