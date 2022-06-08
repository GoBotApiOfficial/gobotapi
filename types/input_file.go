package types

import (
	"github.com/Squirrel-Network/gobotapi/types/raw"
	"os"
	"path"
)

func InputFile(filePath string) raw.InputFile {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil
	}
	return InputBytes{
		Name: path.Base(filePath),
		Data: file,
	}
}
