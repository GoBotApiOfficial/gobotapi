// Code AutoGenerated; DO NOT EDIT.

package types

import (
	"encoding/json"
	"fmt"
)

// InlineQueryResultPhoto Represents a link to a photo
// By default, this photo will be sent by the user with optional caption
// Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
	Caption               string                `json:"caption,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	Description           string                `json:"description,omitempty"`
	ID                    string                `json:"id"`
	InputMessageContent   any                   `json:"input_message_content,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	PhotoHeight           int                   `json:"photo_height,omitempty"`
	PhotoURL              string                `json:"photo_url"`
	PhotoWidth            int64                 `json:"photo_width,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	Title                 string                `json:"title,omitempty"`
}

func (entity InlineQueryResultPhoto) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type                  string                `json:"type"`
		ID                    string                `json:"id"`
		PhotoURL              string                `json:"photo_url"`
		ThumbnailURL          string                `json:"thumbnail_url"`
		PhotoWidth            int64                 `json:"photo_width,omitempty"`
		PhotoHeight           int                   `json:"photo_height,omitempty"`
		Title                 string                `json:"title,omitempty"`
		Description           string                `json:"description,omitempty"`
		Caption               string                `json:"caption,omitempty"`
		ParseMode             string                `json:"parse_mode,omitempty"`
		CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
		ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
		ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent   any                   `json:"input_message_content,omitempty"`
	}{
		Type:                  "photo",
		ID:                    entity.ID,
		PhotoURL:              entity.PhotoURL,
		ThumbnailURL:          entity.ThumbnailURL,
		PhotoWidth:            entity.PhotoWidth,
		PhotoHeight:           entity.PhotoHeight,
		Title:                 entity.Title,
		Description:           entity.Description,
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
