package types

import "encoding/json"
import "fmt"

type InlineQueryResultPhoto struct {
	Caption string `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Description string `json:"description,omitempty"`
	Id string `json:"id"`
	InputMessageContent interface{} `json:"input_message_content,omitempty"`
	ParseMode string `json:"parse_mode,omitempty"`
	PhotoHeight int `json:"photo_height,omitempty"`
	PhotoUrl string `json:"photo_url"`
	PhotoWidth int64 `json:"photo_width,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbUrl string `json:"thumb_url"`
	Title string `json:"title,omitempty"`
}

func (entity InlineQueryResultPhoto) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type string `json:"type"`
		Id string `json:"id"`
		PhotoUrl string `json:"photo_url"`
		ThumbUrl string `json:"thumb_url"`
		PhotoWidth int64 `json:"photo_width,omitempty"`
		PhotoHeight int `json:"photo_height,omitempty"`
		Title string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		Caption string `json:"caption,omitempty"`
		ParseMode string `json:"parse_mode,omitempty"`
		CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
	} {
		Type: "photo",
		Id: entity.Id,
		PhotoUrl: entity.PhotoUrl,
		ThumbUrl: entity.ThumbUrl,
		PhotoWidth: entity.PhotoWidth,
		PhotoHeight: entity.PhotoHeight,
		Title: entity.Title,
		Description: entity.Description,
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
