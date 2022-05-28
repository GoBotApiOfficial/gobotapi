package types

import "encoding/json"
import "fmt"

// InlineQueryResultMpeg4Gif Represents a link to a video animation (H.264/MPEG-4 AVC video without sound)
// By default, this animated MPEG-4 file will be sent by the user with optional caption
// Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultMpeg4Gif struct {
	Caption string `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Id string `json:"id"`
	InputMessageContent interface{} `json:"input_message_content,omitempty"`
	Mpeg4Duration int `json:"mpeg4_duration,omitempty"`
	Mpeg4Height int `json:"mpeg4_height,omitempty"`
	Mpeg4Url string `json:"mpeg4_url"`
	Mpeg4Width int64 `json:"mpeg4_width,omitempty"`
	ParseMode string `json:"parse_mode,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbMimeType string `json:"thumb_mime_type,omitempty"`
	ThumbUrl string `json:"thumb_url"`
	Title string `json:"title,omitempty"`
}

func (entity InlineQueryResultMpeg4Gif) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type string `json:"type"`
		Id string `json:"id"`
		Mpeg4Url string `json:"mpeg4_url"`
		Mpeg4Width int64 `json:"mpeg4_width,omitempty"`
		Mpeg4Height int `json:"mpeg4_height,omitempty"`
		Mpeg4Duration int `json:"mpeg4_duration,omitempty"`
		ThumbUrl string `json:"thumb_url"`
		ThumbMimeType string `json:"thumb_mime_type,omitempty"`
		Title string `json:"title,omitempty"`
		Caption string `json:"caption,omitempty"`
		ParseMode string `json:"parse_mode,omitempty"`
		CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
	} {
		Type: "mpeg4_gif",
		Id: entity.Id,
		Mpeg4Url: entity.Mpeg4Url,
		Mpeg4Width: entity.Mpeg4Width,
		Mpeg4Height: entity.Mpeg4Height,
		Mpeg4Duration: entity.Mpeg4Duration,
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
