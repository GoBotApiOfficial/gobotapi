package types

type Audio struct {
	Duration int `json:"duration"`
	FileId string `json:"file_id"`
	FileName string `json:"file_name,omitempty"`
	FileSize int `json:"file_size,omitempty"`
	FileUniqueId string `json:"file_unique_id"`
	MimeType string `json:"mime_type,omitempty"`
	Performer string `json:"performer,omitempty"`
	Thumb *PhotoSize `json:"thumb,omitempty"`
	Title string `json:"title,omitempty"`
}
