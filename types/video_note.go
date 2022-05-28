package types

// VideoNote Represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	Duration int `json:"duration"`
	FileId string `json:"file_id"`
	FileSize int `json:"file_size,omitempty"`
	FileUniqueId string `json:"file_unique_id"`
	Length int `json:"length"`
	Thumb *PhotoSize `json:"thumb,omitempty"`
}
