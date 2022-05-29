package gobotapi

import (
	"errors"
	"os"
	"path"
)

func (ctx *Client) DownloadFile(fileId, filePath string) error {
	bytes, err := ctx.DownloadBytes(fileId)
	if err != nil {
		return err
	}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return errors.New("path not found")
	} else {
		res, _ := os.Stat(path.Dir(filePath))
		_ = os.WriteFile(filePath, bytes, res.Mode().Perm())
	}
	return nil
}
