package types

import "encoding/json"
import "fmt"

// InlineQueryResultDocument Represents a link to a file
// By default, this file will be sent by the user with an optional caption
// Alternatively, you can use input_message_content to send a message with the specified content instead of the file
// Currently, only .PDF and .ZIP files can be sent using this method.
// Note: This will only work in Telegram versions released after 9 April, 2016
// Older clients will ignore them.
type InlineQueryResultDocument struct {
	Caption string `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Description string `json:"description,omitempty"`
	DocumentUrl string `json:"document_url"`
	Id string `json:"id"`
	InputMessageContent interface{} `json:"input_message_content,omitempty"`
	MimeType string `json:"mime_type"`
	ParseMode string `json:"parse_mode,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbHeight int `json:"thumb_height,omitempty"`
	ThumbUrl string `json:"thumb_url,omitempty"`
	ThumbWidth int64 `json:"thumb_width,omitempty"`
	Title string `json:"title"`
}

func (entity InlineQueryResultDocument) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type string `json:"type"`
		Id string `json:"id"`
		Title string `json:"title"`
		Caption string `json:"caption,omitempty"`
		ParseMode string `json:"parse_mode,omitempty"`
		CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
		DocumentUrl string `json:"document_url"`
		MimeType string `json:"mime_type"`
		Description string `json:"description,omitempty"`
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
		ThumbUrl string `json:"thumb_url,omitempty"`
		ThumbWidth int64 `json:"thumb_width,omitempty"`
		ThumbHeight int `json:"thumb_height,omitempty"`
	} {
		Type: "document",
		Id: entity.Id,
		Title: entity.Title,
		Caption: entity.Caption,
		ParseMode: entity.ParseMode,
		CaptionEntities: entity.CaptionEntities,
		DocumentUrl: entity.DocumentUrl,
		MimeType: entity.MimeType,
		Description: entity.Description,
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
