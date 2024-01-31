// Code AutoGenerated; DO NOT EDIT.

package types

import (
	"encoding/json"
	"fmt"
)

// InlineQueryResultLocation Represents a location on a map
// By default, the location will be sent by the user
// Alternatively, you can use input_message_content to send a message with the specified content instead of the location.
type InlineQueryResultLocation struct {
	Heading              int                   `json:"heading,omitempty"`
	HorizontalAccuracy   float64               `json:"horizontal_accuracy,omitempty"`
	ID                   string                `json:"id"`
	InputMessageContent  any                   `json:"input_message_content,omitempty"`
	Latitude             float64               `json:"latitude"`
	LivePeriod           int                   `json:"live_period,omitempty"`
	Longitude            float64               `json:"longitude"`
	ProximityAlertRadius int                   `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbnailHeight      int                   `json:"thumbnail_height,omitempty"`
	ThumbnailURL         string                `json:"thumbnail_url,omitempty"`
	ThumbnailWidth       int64                 `json:"thumbnail_width,omitempty"`
	Title                string                `json:"title"`
}

func (entity InlineQueryResultLocation) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type                 string                `json:"type"`
		ID                   string                `json:"id"`
		Latitude             float64               `json:"latitude"`
		Longitude            float64               `json:"longitude"`
		Title                string                `json:"title"`
		HorizontalAccuracy   float64               `json:"horizontal_accuracy,omitempty"`
		LivePeriod           int                   `json:"live_period,omitempty"`
		Heading              int                   `json:"heading,omitempty"`
		ProximityAlertRadius int                   `json:"proximity_alert_radius,omitempty"`
		ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent  any                   `json:"input_message_content,omitempty"`
		ThumbnailURL         string                `json:"thumbnail_url,omitempty"`
		ThumbnailWidth       int64                 `json:"thumbnail_width,omitempty"`
		ThumbnailHeight      int                   `json:"thumbnail_height,omitempty"`
	}{
		Type:                 "location",
		ID:                   entity.ID,
		Latitude:             entity.Latitude,
		Longitude:            entity.Longitude,
		Title:                entity.Title,
		HorizontalAccuracy:   entity.HorizontalAccuracy,
		LivePeriod:           entity.LivePeriod,
		Heading:              entity.Heading,
		ProximityAlertRadius: entity.ProximityAlertRadius,
		ReplyMarkup:          entity.ReplyMarkup,
		InputMessageContent:  entity.InputMessageContent,
		ThumbnailURL:         entity.ThumbnailURL,
		ThumbnailWidth:       entity.ThumbnailWidth,
		ThumbnailHeight:      entity.ThumbnailHeight,
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
