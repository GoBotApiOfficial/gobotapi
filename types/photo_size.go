package types

type PhotoSize struct {
	FileId string `json:"file_id"`
	FileSize int `json:"file_size,omitempty"`
	FileUniqueId string `json:"file_unique_id"`
	Height int `json:"height"`
	Width int64 `json:"width"`
}