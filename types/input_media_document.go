package types

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "fmt"
import "encoding/json"

// InputMediaDocument Represents a general file to be sent.
type InputMediaDocument struct {
	Caption string `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
	Media rawTypes.InputFile `json:"media,omitempty"`
	ParseMode string `json:"parse_mode,omitempty"`
	Thumb rawTypes.InputFile `json:"thumb,omitempty"`
}

func (entity *InputMediaDocument) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Media.(type) {
		case InputFile:
			files["document"] = entity.Media
	}
	switch entity.Thumb.(type) {
		case InputFile:
			files["thumb"] = entity.Thumb
	}
	return files
}

func (entity *InputMediaDocument) SetAttachment(attach string) {
	entity.Media = InputURL(fmt.Sprintf("attach://%s", attach))
}

func (entity *InputMediaDocument) SetAttachmentThumb(attach string) {
	entity.Thumb = InputURL(fmt.Sprintf("attach://%s", attach))
}

func (entity InputMediaDocument) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type string `json:"type"`
		Media rawTypes.InputFile `json:"media,omitempty"`
		Thumb rawTypes.InputFile `json:"thumb,omitempty"`
		Caption string `json:"caption,omitempty"`
		ParseMode string `json:"parse_mode,omitempty"`
		CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
		DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
	} {
		Type: "document",
		Media: entity.Media,
		Thumb: entity.Thumb,
		Caption: entity.Caption,
		ParseMode: entity.ParseMode,
		CaptionEntities: entity.CaptionEntities,
		DisableContentTypeDetection: entity.DisableContentTypeDetection,
	}
	return json.Marshal(alias)
}
