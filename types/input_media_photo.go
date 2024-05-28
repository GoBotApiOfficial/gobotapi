// Code AutoGenerated; DO NOT EDIT.

package types

import (
	"encoding/json"
	"fmt"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// InputMediaPhoto Represents a photo to be sent.
type InputMediaPhoto struct {
	Caption               string             `json:"caption,omitempty"`
	CaptionEntities       []MessageEntity    `json:"caption_entities,omitempty"`
	HasSpoiler            bool               `json:"has_spoiler,omitempty"`
	Media                 rawTypes.InputFile `json:"media,omitempty"`
	ParseMode             string             `json:"parse_mode,omitempty"`
	ShowCaptionAboveMedia bool               `json:"show_caption_above_media,omitempty"`
}

func (entity *InputMediaPhoto) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Media.(type) {
	case InputBytes:
		files["photo"] = entity.Media
	}
	return files
}

func (entity *InputMediaPhoto) SetAttachment(attach string) {
	entity.Media = InputURL(fmt.Sprintf("attach://%s", attach))
}

func (entity *InputMediaPhoto) SetAttachmentThumb(_ string) {
}

func (entity InputMediaPhoto) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type                  string             `json:"type"`
		Media                 rawTypes.InputFile `json:"media,omitempty"`
		Caption               string             `json:"caption,omitempty"`
		ParseMode             string             `json:"parse_mode,omitempty"`
		CaptionEntities       []MessageEntity    `json:"caption_entities,omitempty"`
		ShowCaptionAboveMedia bool               `json:"show_caption_above_media,omitempty"`
		HasSpoiler            bool               `json:"has_spoiler,omitempty"`
	}{
		Type:                  "photo",
		Media:                 entity.Media,
		Caption:               entity.Caption,
		ParseMode:             entity.ParseMode,
		CaptionEntities:       entity.CaptionEntities,
		ShowCaptionAboveMedia: entity.ShowCaptionAboveMedia,
		HasSpoiler:            entity.HasSpoiler,
	}
	return json.Marshal(alias)
}
