package types

import "encoding/json"
import "fmt"

type InlineQueryResultAudio struct {
	AudioDuration int `json:"audio_duration,omitempty"`
	AudioUrl string `json:"audio_url"`
	Caption string `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Id string `json:"id"`
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
	ParseMode string `json:"parse_mode,omitempty"`
	Performer string `json:"performer,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Title string `json:"title"`
}

func (entity InlineQueryResultAudio) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type string `json:"type"`
		Id string `json:"id"`
		AudioUrl string `json:"audio_url"`
		Title string `json:"title"`
		Caption string `json:"caption,omitempty"`
		ParseMode string `json:"parse_mode,omitempty"`
		CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
		Performer string `json:"performer,omitempty"`
		AudioDuration int `json:"audio_duration,omitempty"`
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
	} {
		Type: "audio",
		Id: entity.Id,
		AudioUrl: entity.AudioUrl,
		Title: entity.Title,
		Caption: entity.Caption,
		ParseMode: entity.ParseMode,
		CaptionEntities: entity.CaptionEntities,
		Performer: entity.Performer,
		AudioDuration: entity.AudioDuration,
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
