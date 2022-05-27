package types

type File struct {
	FileId string `json:"file_id"`
	FilePath string `json:"file_path,omitempty"`
	FileSize int `json:"file_size,omitempty"`
	FileUniqueId string `json:"file_unique_id"`
}
