package types

import "encoding/json"
import "fmt"

type InlineQueryResultVideo struct {
	Caption string `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Description string `json:"description,omitempty"`
	Id string `json:"id"`
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
	MimeType string `json:"mime_type"`
	ParseMode string `json:"parse_mode,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbUrl string `json:"thumb_url"`
	Title string `json:"title"`
	VideoDuration int64 `json:"video_duration,omitempty"`
	VideoHeight int64 `json:"video_height,omitempty"`
	VideoUrl string `json:"video_url"`
	VideoWidth int64 `json:"video_width,omitempty"`
}

func (entity InlineQueryResultVideo) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type string `json:"type"`
		Id string `json:"id"`
		VideoUrl string `json:"video_url"`
		MimeType string `json:"mime_type"`
		ThumbUrl string `json:"thumb_url"`
		Title string `json:"title"`
		Caption string `json:"caption,omitempty"`
		ParseMode string `json:"parse_mode,omitempty"`
		CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
		VideoWidth int64 `json:"video_width,omitempty"`
		VideoHeight int64 `json:"video_height,omitempty"`
		VideoDuration int64 `json:"video_duration,omitempty"`
		Description string `json:"description,omitempty"`
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
	} {
		Type: "video",
		Id: entity.Id,
		VideoUrl: entity.VideoUrl,
		MimeType: entity.MimeType,
		ThumbUrl: entity.ThumbUrl,
		Title: entity.Title,
		Caption: entity.Caption,
		ParseMode: entity.ParseMode,
		CaptionEntities: entity.CaptionEntities,
		VideoWidth: entity.VideoWidth,
		VideoHeight: entity.VideoHeight,
		VideoDuration: entity.VideoDuration,
		Description: entity.Description,
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
