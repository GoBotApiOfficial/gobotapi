package types

import "encoding/json"
import "fmt"

type InlineQueryResultCachedVoice struct {
	Caption string `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Id string `json:"id"`
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
	ParseMode string `json:"parse_mode,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Title string `json:"title"`
	VoiceFileId string `json:"voice_file_id"`
}

func (entity InlineQueryResultCachedVoice) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type string `json:"type"`
		Id string `json:"id"`
		VoiceFileId string `json:"voice_file_id"`
		Title string `json:"title"`
		Caption string `json:"caption,omitempty"`
		ParseMode string `json:"parse_mode,omitempty"`
		CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
	} {
		Type: "voice",
		Id: entity.Id,
		VoiceFileId: entity.VoiceFileId,
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
