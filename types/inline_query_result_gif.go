// Code AutoGenerated; DO NOT EDIT.

package types

import (
	"encoding/json"
	"fmt"
)

// InlineQueryResultGif Represents a link to an animated GIF file
// By default, this animated GIF file will be sent by the user with optional caption
// Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultGif struct {
	Caption               string                `json:"caption,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	GifDuration           int                   `json:"gif_duration,omitempty"`
	GifHeight             int                   `json:"gif_height,omitempty"`
	GifURL                string                `json:"gif_url"`
	GifWidth              int64                 `json:"gif_width,omitempty"`
	ID                    string                `json:"id"`
	InputMessageContent   any                   `json:"input_message_content,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	ThumbnailMimeType     string                `json:"thumbnail_mime_type,omitempty"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	Title                 string                `json:"title,omitempty"`
}

func (entity InlineQueryResultGif) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type                  string                `json:"type"`
		ID                    string                `json:"id"`
		GifURL                string                `json:"gif_url"`
		GifWidth              int64                 `json:"gif_width,omitempty"`
		GifHeight             int                   `json:"gif_height,omitempty"`
		GifDuration           int                   `json:"gif_duration,omitempty"`
		ThumbnailURL          string                `json:"thumbnail_url"`
		ThumbnailMimeType     string                `json:"thumbnail_mime_type,omitempty"`
		Title                 string                `json:"title,omitempty"`
		Caption               string                `json:"caption,omitempty"`
		ParseMode             string                `json:"parse_mode,omitempty"`
		CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
		ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
		ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent   any                   `json:"input_message_content,omitempty"`
	}{
		Type:                  "gif",
		ID:                    entity.ID,
		GifURL:                entity.GifURL,
		GifWidth:              entity.GifWidth,
		GifHeight:             entity.GifHeight,
		GifDuration:           entity.GifDuration,
		ThumbnailURL:          entity.ThumbnailURL,
		ThumbnailMimeType:     entity.ThumbnailMimeType,
		Title:                 entity.Title,
		Caption:               entity.Caption,
		ParseMode:             entity.ParseMode,
		CaptionEntities:       entity.CaptionEntities,
		ShowCaptionAboveMedia: entity.ShowCaptionAboveMedia,
		ReplyMarkup:           entity.ReplyMarkup,
		InputMessageContent:   entity.InputMessageContent,
	}
	if entity.InputMessageContent != nil {
		switch entity.InputMessageContent.(type) {
		case InputTextMessageContent, InputLocationMessageContent, InputVenueMessageContent, InputContactMessageContent, InputInvoiceMessageContent:
			break
		default:
			return nil, fmt.Errorf("input_message_content: unknown type: %T", entity.InputMessageContent)
		}
	}
	return json.Marshal(alias)
}
