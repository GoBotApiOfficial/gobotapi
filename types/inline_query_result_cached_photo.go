package types

import "encoding/json"
import "fmt"

type InlineQueryResultCachedPhoto struct {
	Caption string `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Description string `json:"description,omitempty"`
	Id string `json:"id"`
	InputMessageContent interface{} `json:"input_message_content"`
	ParseMode string `json:"parse_mode,omitempty"`
	PhotoFileId string `json:"photo_file_id"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Title string `json:"title,omitempty"`
}

func (entity InlineQueryResultCachedPhoto) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type string `json:"type"`
		Id string `json:"id"`
		PhotoFileId string `json:"photo_file_id"`
		Title string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		Caption string `json:"caption,omitempty"`
		ParseMode string `json:"parse_mode,omitempty"`
		CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent interface{} `json:"input_message_content"`
	} {
		Type: "photo",
		Id: entity.Id,
		PhotoFileId: entity.PhotoFileId,
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
