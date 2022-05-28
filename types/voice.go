package types

// Voice Represents a voice note.
type Voice struct {
	Duration int `json:"duration"`
	FileId string `json:"file_id"`
	FileSize int `json:"file_size,omitempty"`
	FileUniqueId string `json:"file_unique_id"`
	MimeType string `json:"mime_type,omitempty"`
}
