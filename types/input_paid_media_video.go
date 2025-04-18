// Code AutoGenerated; DO NOT EDIT.

package types

import (
	"encoding/json"
	"fmt"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// InputPaidMediaVideo The paid media to send is a video.
type InputPaidMediaVideo struct {
	Cover             string             `json:"cover,omitempty"`
	Duration          int                `json:"duration,omitempty"`
	Height            int                `json:"height,omitempty"`
	Media             rawTypes.InputFile `json:"media,omitempty"`
	StartTimestamp    int                `json:"start_timestamp,omitempty"`
	SupportsStreaming bool               `json:"supports_streaming,omitempty"`
	Thumbnail         rawTypes.InputFile `json:"thumbnail,omitempty,omitempty"`
	Width             int64              `json:"width,omitempty"`
}

func (entity *InputPaidMediaVideo) Files() map[string]rawTypes.InputFile {
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

func (entity *InputPaidMediaVideo) SetAttachment(attach string) {
	entity.Media = InputURL(fmt.Sprintf("attach://%s", attach))
}

func (entity *InputPaidMediaVideo) SetAttachmentThumb(attach string) {
	entity.Thumbnail = InputURL(fmt.Sprintf("attach://%s", attach))
}

func (entity InputPaidMediaVideo) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type              string             `json:"type"`
		Media             rawTypes.InputFile `json:"media,omitempty"`
		Thumbnail         rawTypes.InputFile `json:"thumbnail,omitempty"`
		Cover             string             `json:"cover,omitempty"`
		StartTimestamp    int                `json:"start_timestamp,omitempty"`
		Width             int64              `json:"width,omitempty"`
		Height            int                `json:"height,omitempty"`
		Duration          int                `json:"duration,omitempty"`
		SupportsStreaming bool               `json:"supports_streaming,omitempty"`
	}{
		Type:              "video",
		Media:             entity.Media,
		Thumbnail:         entity.Thumbnail,
		Cover:             entity.Cover,
		StartTimestamp:    entity.StartTimestamp,
		Width:             entity.Width,
		Height:            entity.Height,
		Duration:          entity.Duration,
		SupportsStreaming: entity.SupportsStreaming,
	}
	return json.Marshal(alias)
}
