package types

import "encoding/json"
import "fmt"

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
