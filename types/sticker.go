package types

type Sticker struct {
	Emoji string `json:"emoji,omitempty"`
	FileId string `json:"file_id"`
	FileSize int `json:"file_size,omitempty"`
	FileUniqueId string `json:"file_unique_id"`
	Height int `json:"height"`
	IsAnimated bool `json:"is_animated"`
	IsVideo bool `json:"is_video"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	SetName string `json:"set_name,omitempty"`
	Thumb *PhotoSize `json:"thumb,omitempty"`
	Width int64 `json:"width"`
}
