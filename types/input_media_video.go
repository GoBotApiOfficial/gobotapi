package types

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "fmt"
import "encoding/json"

type InputMediaVideo struct {
	Caption string `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Duration int `json:"duration,omitempty"`
	Height int `json:"height,omitempty"`
	Media rawTypes.InputFile `json:"media,omitempty"`
	ParseMode string `json:"parse_mode,omitempty"`
	SupportsStreaming bool `json:"supports_streaming,omitempty"`
	Thumb rawTypes.InputFile `json:"thumb,omitempty"`
	Width int64 `json:"width,omitempty"`
}

func (entity *InputMediaVideo) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Media.(type) {
		case InputFile:
			files["video"] = entity.Media
	}
	switch entity.Thumb.(type) {
		case InputFile:
			files["thumb"] = entity.Thumb
	}
	return files
}

func (entity *InputMediaVideo) SetAttachment(attach string) {
	entity.Media = InputURL(fmt.Sprintf("attach://%s", attach))
}

func (entity *InputMediaVideo) SetAttachmentThumb(attach string) {
	entity.Thumb = InputURL(fmt.Sprintf("attach://%s", attach))
}

func (entity InputMediaVideo) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type string `json:"type"`
		Media rawTypes.InputFile `json:"media,omitempty"`
		Thumb rawTypes.InputFile `json:"thumb,omitempty"`
		Caption string `json:"caption,omitempty"`
		ParseMode string `json:"parse_mode,omitempty"`
		CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
		Width int64 `json:"width,omitempty"`
		Height int `json:"height,omitempty"`
		Duration int `json:"duration,omitempty"`
		SupportsStreaming bool `json:"supports_streaming,omitempty"`
	} {
		Type: "video",
		Media: entity.Media,
		Thumb: entity.Thumb,
		Caption: entity.Caption,
		ParseMode: entity.ParseMode,
		CaptionEntities: entity.CaptionEntities,
		Width: entity.Width,
		Height: entity.Height,
		Duration: entity.Duration,
		SupportsStreaming: entity.SupportsStreaming,
	}
	return json.Marshal(alias)
}