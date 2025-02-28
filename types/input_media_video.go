// Code AutoGenerated; DO NOT EDIT.

package types

import (
	"encoding/json"
	"fmt"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// InputMediaVideo Represents a video to be sent.
type InputMediaVideo struct {
	Caption               string             `json:"caption,omitempty"`
	CaptionEntities       []MessageEntity    `json:"caption_entities,omitempty"`
	Cover                 string             `json:"cover,omitempty"`
	Duration              int                `json:"duration,omitempty"`
	HasSpoiler            bool               `json:"has_spoiler,omitempty"`
	Height                int                `json:"height,omitempty"`
	Media                 rawTypes.InputFile `json:"media,omitempty"`
	ParseMode             string             `json:"parse_mode,omitempty"`
	ShowCaptionAboveMedia bool               `json:"show_caption_above_media,omitempty"`
	StartTimestamp        int                `json:"start_timestamp,omitempty"`
	SupportsStreaming     bool               `json:"supports_streaming,omitempty"`
	Thumbnail             rawTypes.InputFile `json:"thumbnail,omitempty,omitempty"`
	Width                 int64              `json:"width,omitempty"`
}

func (entity *InputMediaVideo) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Media.(type) {
	case InputBytes:
		files["video"] = entity.Media
	}
	switch entity.Thumbnail.(type) {
	case InputBytes:
		files["thumbnail"] = entity.Thumbnail
	}
	return files
}

func (entity *InputMediaVideo) SetAttachment(attach string) {
	entity.Media = InputURL(fmt.Sprintf("attach://%s", attach))
}

func (entity *InputMediaVideo) SetAttachmentThumb(attach string) {
	entity.Thumbnail = InputURL(fmt.Sprintf("attach://%s", attach))
}

func (entity InputMediaVideo) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type                  string             `json:"type"`
		Media                 rawTypes.InputFile `json:"media,omitempty"`
		Thumbnail             rawTypes.InputFile `json:"thumbnail,omitempty"`
		Cover                 string             `json:"cover,omitempty"`
		StartTimestamp        int                `json:"start_timestamp,omitempty"`
		Caption               string             `json:"caption,omitempty"`
		ParseMode             string             `json:"parse_mode,omitempty"`
		CaptionEntities       []MessageEntity    `json:"caption_entities,omitempty"`
		ShowCaptionAboveMedia bool               `json:"show_caption_above_media,omitempty"`
		Width                 int64              `json:"width,omitempty"`
		Height                int                `json:"height,omitempty"`
		Duration              int                `json:"duration,omitempty"`
		SupportsStreaming     bool               `json:"supports_streaming,omitempty"`
		HasSpoiler            bool               `json:"has_spoiler,omitempty"`
	}{
		Type:                  "video",
		Media:                 entity.Media,
		Thumbnail:             entity.Thumbnail,
		Cover:                 entity.Cover,
		StartTimestamp:        entity.StartTimestamp,
		Caption:               entity.Caption,
		ParseMode:             entity.ParseMode,
		CaptionEntities:       entity.CaptionEntities,
		ShowCaptionAboveMedia: entity.ShowCaptionAboveMedia,
		Width:                 entity.Width,
		Height:                entity.Height,
		Duration:              entity.Duration,
		SupportsStreaming:     entity.SupportsStreaming,
		HasSpoiler:            entity.HasSpoiler,
	}
	return json.Marshal(alias)
}
