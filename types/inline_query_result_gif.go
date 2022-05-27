package types

import "encoding/json"
import "fmt"

type InlineQueryResultGif struct {
	Caption string `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	GifDuration int `json:"gif_duration,omitempty"`
	GifHeight int `json:"gif_height,omitempty"`
	GifUrl string `json:"gif_url"`
	GifWidth int64 `json:"gif_width,omitempty"`
	Id string `json:"id"`
	InputMessageContent interface{} `json:"input_message_content"`
	ParseMode string `json:"parse_mode,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbMimeType string `json:"thumb_mime_type,omitempty"`
	ThumbUrl string `json:"thumb_url"`
	Title string `json:"title,omitempty"`
}

func (entity InlineQueryResultGif) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type string `json:"type"`
		Id string `json:"id"`
		GifUrl string `json:"gif_url"`
		GifWidth int64 `json:"gif_width,omitempty"`
		GifHeight int `json:"gif_height,omitempty"`
		GifDuration int `json:"gif_duration,omitempty"`
		ThumbUrl string `json:"thumb_url"`
		ThumbMimeType string `json:"thumb_mime_type,omitempty"`
		Title string `json:"title,omitempty"`
		Caption string `json:"caption,omitempty"`
		ParseMode string `json:"parse_mode,omitempty"`
		CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent interface{} `json:"input_message_content"`
	} {
		Type: "gif",
		Id: entity.Id,
		GifUrl: entity.GifUrl,
		GifWidth: entity.GifWidth,
		GifHeight: entity.GifHeight,
		GifDuration: entity.GifDuration,
		ThumbUrl: entity.ThumbUrl,
		ThumbMimeType: entity.ThumbMimeType,
		Title: entity.Title,
		Caption: entity.Caption,
		ParseMode: entity.ParseMode,
		CaptionEntities: entity.CaptionEntities,
		ReplyMarkup: entity.ReplyMarkup,
		InputMessageContent: entity.InputMessageContent,
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
