package types

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "fmt"
import "encoding/json"

// InputMediaPhoto Represents a photo to be sent.
type InputMediaPhoto struct {
	Caption string `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Media rawTypes.InputFile `json:"media,omitempty"`
	ParseMode string `json:"parse_mode,omitempty"`
}

func (entity *InputMediaPhoto) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Media.(type) {
		case InputFile:
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
		Type string `json:"type"`
		Media rawTypes.InputFile `json:"media,omitempty"`
		Caption string `json:"caption,omitempty"`
		ParseMode string `json:"parse_mode,omitempty"`
		CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	} {
		Type: "photo",
		Media: entity.Media,
		Caption: entity.Caption,
		ParseMode: entity.ParseMode,
		CaptionEntities: entity.CaptionEntities,
	}
	return json.Marshal(alias)
}
