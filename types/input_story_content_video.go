// Code AutoGenerated; DO NOT EDIT.

package types

import "encoding/json"

// InputStoryContentVideo Describes a video to post as a story.
type InputStoryContentVideo struct {
	CoverFrameTimestamp float64 `json:"cover_frame_timestamp,omitempty"`
	Duration            float64 `json:"duration,omitempty"`
	IsAnimation         bool    `json:"is_animation,omitempty"`
	Video               string  `json:"video"`
}

func (entity InputStoryContentVideo) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type                string  `json:"type"`
		Video               string  `json:"video"`
		Duration            float64 `json:"duration,omitempty"`
		CoverFrameTimestamp float64 `json:"cover_frame_timestamp,omitempty"`
		IsAnimation         bool    `json:"is_animation,omitempty"`
	}{
		Type:                "“video”",
		Video:               entity.Video,
		Duration:            entity.Duration,
		CoverFrameTimestamp: entity.CoverFrameTimestamp,
		IsAnimation:         entity.IsAnimation,
	}
	return json.Marshal(alias)
}
