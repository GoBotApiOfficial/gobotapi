package types

import "encoding/json"
import "fmt"

// InlineQueryResultVenue Represents a venue
// By default, the venue will be sent by the user
// Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.
// Note: This will only work in Telegram versions released after 9 April, 2016
// Older clients will ignore them.
type InlineQueryResultVenue struct {
	Address string `json:"address"`
	FoursquareId string `json:"foursquare_id,omitempty"`
	FoursquareType string `json:"foursquare_type,omitempty"`
	GooglePlaceId string `json:"google_place_id,omitempty"`
	GooglePlaceType string `json:"google_place_type,omitempty"`
	Id string `json:"id"`
	InputMessageContent interface{} `json:"input_message_content,omitempty"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbHeight int `json:"thumb_height,omitempty"`
	ThumbUrl string `json:"thumb_url,omitempty"`
	ThumbWidth int64 `json:"thumb_width,omitempty"`
	Title string `json:"title"`
}

func (entity InlineQueryResultVenue) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type string `json:"type"`
		Id string `json:"id"`
		Latitude float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Title string `json:"title"`
		Address string `json:"address"`
		FoursquareId string `json:"foursquare_id,omitempty"`
		FoursquareType string `json:"foursquare_type,omitempty"`
		GooglePlaceId string `json:"google_place_id,omitempty"`
		GooglePlaceType string `json:"google_place_type,omitempty"`
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
		ThumbUrl string `json:"thumb_url,omitempty"`
		ThumbWidth int64 `json:"thumb_width,omitempty"`
		ThumbHeight int `json:"thumb_height,omitempty"`
	} {
		Type: "venue",
		Id: entity.Id,
		Latitude: entity.Latitude,
		Longitude: entity.Longitude,
		Title: entity.Title,
		Address: entity.Address,
		FoursquareId: entity.FoursquareId,
		FoursquareType: entity.FoursquareType,
		GooglePlaceId: entity.GooglePlaceId,
		GooglePlaceType: entity.GooglePlaceType,
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
