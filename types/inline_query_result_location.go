package types

import "encoding/json"
import "fmt"

type InlineQueryResultLocation struct {
	Heading int `json:"heading,omitempty"`
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
	Id string `json:"id"`
	InputMessageContent interface{} `json:"input_message_content,omitempty"`
	Latitude float64 `json:"latitude"`
	LivePeriod int `json:"live_period,omitempty"`
	Longitude float64 `json:"longitude"`
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbHeight int `json:"thumb_height,omitempty"`
	ThumbUrl string `json:"thumb_url,omitempty"`
	ThumbWidth int64 `json:"thumb_width,omitempty"`
	Title string `json:"title"`
}

func (entity InlineQueryResultLocation) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type string `json:"type"`
		Id string `json:"id"`
		Latitude float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Title string `json:"title"`
		HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
		LivePeriod int `json:"live_period,omitempty"`
		Heading int `json:"heading,omitempty"`
		ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
		ThumbUrl string `json:"thumb_url,omitempty"`
		ThumbWidth int64 `json:"thumb_width,omitempty"`
		ThumbHeight int `json:"thumb_height,omitempty"`
	} {
		Type: "location",
		Id: entity.Id,
		Latitude: entity.Latitude,
		Longitude: entity.Longitude,
		Title: entity.Title,
		HorizontalAccuracy: entity.HorizontalAccuracy,
		LivePeriod: entity.LivePeriod,
		Heading: entity.Heading,
		ProximityAlertRadius: entity.ProximityAlertRadius,
		ReplyMarkup: entity.ReplyMarkup,
		InputMessageContent: entity.InputMessageContent,
		ThumbUrl: entity.ThumbUrl,
		ThumbWidth: entity.ThumbWidth,
		ThumbHeight: entity.ThumbHeight,
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
