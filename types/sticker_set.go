package types

type StickerSet struct {
	ContainsMasks bool `json:"contains_masks"`
	IsAnimated bool `json:"is_animated"`
	IsVideo bool `json:"is_video"`
	Name string `json:"name"`
	Stickers []Sticker `json:"stickers,omitempty"`
	Thumb *PhotoSize `json:"thumb,omitempty"`
	Title string `json:"title"`
}