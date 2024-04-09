// Code AutoGenerated; DO NOT EDIT.

package types

// StickerSet Represents a sticker set.
type StickerSet struct {
	Name        string     `json:"name"`
	StickerType string     `json:"sticker_type"`
	Stickers    []Sticker  `json:"stickers,omitempty"`
	Thumbnail   *PhotoSize `json:"thumbnail,omitempty"`
	Title       string     `json:"title"`
}
