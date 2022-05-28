package types

// File Represents a file ready to be downloaded
// The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>
// It is guaranteed that the link will be valid for at least 1 hour
// When the link expires, a new one can be requested by calling getFile.
type File struct {
	FileId string `json:"file_id"`
	FilePath string `json:"file_path,omitempty"`
	FileSize int `json:"file_size,omitempty"`
	FileUniqueId string `json:"file_unique_id"`
}
