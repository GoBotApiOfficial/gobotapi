package types

import "encoding/json"
import "fmt"

type InlineQueryResultCachedAudio struct {
	AudioFileId string `json:"audio_file_id"`
	Caption string `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Id string `json:"id"`
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
	ParseMode string `json:"parse_mode,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (entity InlineQueryResultCachedAudio) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type string `json:"type"`
		Id string `json:"id"`
		AudioFileId string `json:"audio_file_id"`
		Caption string `json:"caption,omitempty"`
		ParseMode string `json:"parse_mode,omitempty"`
		CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
	} {
		Type: "audio",
		Id: entity.Id,
		AudioFileId: entity.AudioFileId,
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
