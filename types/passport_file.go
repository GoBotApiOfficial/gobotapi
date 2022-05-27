package types

type PassportFile struct {
	FileDate int64 `json:"file_date"`
	FileId string `json:"file_id"`
	FileSize int `json:"file_size"`
	FileUniqueId string `json:"file_unique_id"`
}
