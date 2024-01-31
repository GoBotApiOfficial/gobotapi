// Code AutoGenerated; DO NOT EDIT.

package types

import (
	"encoding/json"
	"fmt"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// InputMediaDocument Represents a general file to be sent.
type InputMediaDocument struct {
	Caption                     string             `json:"caption,omitempty"`
	CaptionEntities             []MessageEntity    `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool               `json:"disable_content_type_detection,omitempty"`
	Media                       rawTypes.InputFile `json:"media,omitempty"`
	ParseMode                   string             `json:"parse_mode,omitempty"`
	Thumbnail                   rawTypes.InputFile `json:"thumbnail,omitempty"`
}

func (entity *InputMediaDocument) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Media.(type) {
	case InputBytes:
		files["document"] = entity.Media
	}
	switch entity.Thumbnail.(type) {
	case InputBytes:
		files["thumbnail"] = entity.Thumbnail
	}
	return files
}

func (entity *InputMediaDocument) SetAttachment(attach string) {
	entity.Media = InputURL(fmt.Sprintf("attach://%s", attach))
}

func (entity *InputMediaDocument) SetAttachmentThumb(_ string) {
}

func (entity InputMediaDocument) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type                        string             `json:"type"`
		Media                       rawTypes.InputFile `json:"media,omitempty"`
		Thumbnail                   rawTypes.InputFile `json:"thumbnail,omitempty"`
		Caption                     string             `json:"caption,omitempty"`
		ParseMode                   string             `json:"parse_mode,omitempty"`
		CaptionEntities             []MessageEntity    `json:"caption_entities,omitempty"`
		DisableContentTypeDetection bool               `json:"disable_content_type_detection,omitempty"`
	}{
		Type:                        "document",
		Media:                       entity.Media,
		Thumbnail:                   entity.Thumbnail,
		Caption:                     entity.Caption,
		ParseMode:                   entity.ParseMode,
		CaptionEntities:             entity.CaptionEntities,
		DisableContentTypeDetection: entity.DisableContentTypeDetection,
	}
	return json.Marshal(alias)
}
