// Code AutoGenerated; DO NOT EDIT.

package types

import (
	"encoding/json"
	"fmt"
)

// InlineQueryResultDocument Represents a link to a file
// By default, this file will be sent by the user with an optional caption
// Alternatively, you can use input_message_content to send a message with the specified content instead of the file
// Currently, only .PDF and .ZIP files can be sent using this method.
type InlineQueryResultDocument struct {
	Caption             string                `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	Description         string                `json:"description,omitempty"`
	DocumentURL         string                `json:"document_url"`
	ID                  string                `json:"id"`
	InputMessageContent any                   `json:"input_message_content,omitempty"`
	MimeType            string                `json:"mime_type"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbnailHeight     int                   `json:"thumbnail_height,omitempty"`
	ThumbnailURL        string                `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int64                 `json:"thumbnail_width,omitempty"`
	Title               string                `json:"title"`
}

func (entity InlineQueryResultDocument) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type                string                `json:"type"`
		ID                  string                `json:"id"`
		Title               string                `json:"title"`
		Caption             string                `json:"caption,omitempty"`
		ParseMode           string                `json:"parse_mode,omitempty"`
		CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
		DocumentURL         string                `json:"document_url"`
		MimeType            string                `json:"mime_type"`
		Description         string                `json:"description,omitempty"`
		ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent any                   `json:"input_message_content,omitempty"`
		ThumbnailURL        string                `json:"thumbnail_url,omitempty"`
		ThumbnailWidth      int64                 `json:"thumbnail_width,omitempty"`
		ThumbnailHeight     int                   `json:"thumbnail_height,omitempty"`
	}{
		Type:                "document",
		ID:                  entity.ID,
		Title:               entity.Title,
		Caption:             entity.Caption,
		ParseMode:           entity.ParseMode,
		CaptionEntities:     entity.CaptionEntities,
		DocumentURL:         entity.DocumentURL,
		MimeType:            entity.MimeType,
		Description:         entity.Description,
		ReplyMarkup:         entity.ReplyMarkup,
		InputMessageContent: entity.InputMessageContent,
		ThumbnailURL:        entity.ThumbnailURL,
		ThumbnailWidth:      entity.ThumbnailWidth,
		ThumbnailHeight:     entity.ThumbnailHeight,
	}
	if !reflect.DeepEqual(entity.InputMessageContent, nil) {
		switch entity.InputMessageContent.(type) {
		case InputTextMessageContent, InputLocationMessageContent, InputVenueMessageContent, InputContactMessageContent, InputInvoiceMessageContent:
			break
		default:
			return nil, fmt.Errorf("input_message_content: unknown type: %T", entity.InputMessageContent)
		}
	}
	return json.Marshal(alias)
}
