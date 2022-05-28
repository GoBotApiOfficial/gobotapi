package types

import "encoding/json"
import "fmt"

// InlineQueryResultContact Represents a contact with a phone number
// By default, this contact will be sent by the user
// Alternatively, you can use input_message_content to send a message with the specified content instead of the contact.
// Note: This will only work in Telegram versions released after 9 April, 2016
// Older clients will ignore them.
type InlineQueryResultContact struct {
	FirstName string `json:"first_name"`
	Id string `json:"id"`
	InputMessageContent interface{} `json:"input_message_content,omitempty"`
	LastName string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbHeight int `json:"thumb_height,omitempty"`
	ThumbUrl string `json:"thumb_url,omitempty"`
	ThumbWidth int64 `json:"thumb_width,omitempty"`
	Vcard string `json:"vcard,omitempty"`
}

func (entity InlineQueryResultContact) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type string `json:"type"`
		Id string `json:"id"`
		PhoneNumber string `json:"phone_number"`
		FirstName string `json:"first_name"`
		LastName string `json:"last_name,omitempty"`
		Vcard string `json:"vcard,omitempty"`
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
		ThumbUrl string `json:"thumb_url,omitempty"`
		ThumbWidth int64 `json:"thumb_width,omitempty"`
		ThumbHeight int `json:"thumb_height,omitempty"`
	} {
		Type: "contact",
		Id: entity.Id,
		PhoneNumber: entity.PhoneNumber,
		FirstName: entity.FirstName,
		LastName: entity.LastName,
		Vcard: entity.Vcard,
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
