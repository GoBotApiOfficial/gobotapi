package types

import "encoding/json"
import "fmt"

type InlineQueryResultCachedMpeg4Gif struct {
	Caption string `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Id string `json:"id"`
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
	Mpeg4FileId string `json:"mpeg4_file_id"`
	ParseMode string `json:"parse_mode,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Title string `json:"title,omitempty"`
}

func (entity InlineQueryResultCachedMpeg4Gif) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type string `json:"type"`
		Id string `json:"id"`
		Mpeg4FileId string `json:"mpeg4_file_id"`
		Title string `json:"title,omitempty"`
		Caption string `json:"caption,omitempty"`
		ParseMode string `json:"parse_mode,omitempty"`
		CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
	} {
		Type: "mpeg4_gif",
		Id: entity.Id,
		Mpeg4FileId: entity.Mpeg4FileId,
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
