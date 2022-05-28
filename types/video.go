package types

// Video Represents a video file.
type Video struct {
	Duration int `json:"duration"`
	FileId string `json:"file_id"`
	FileName string `json:"file_name,omitempty"`
	FileSize int `json:"file_size,omitempty"`
	FileUniqueId string `json:"file_unique_id"`
	Height int `json:"height"`
	MimeType string `json:"mime_type,omitempty"`
	Thumb *PhotoSize `json:"thumb,omitempty"`
	Width int64 `json:"width"`
}
