package types

// Animation Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
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
