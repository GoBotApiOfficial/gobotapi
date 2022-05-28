package types

import "encoding/json"
import "fmt"

// InlineQueryResultCachedVoice Represents a link to a voice message stored on the Telegram servers
// By default, this voice message will be sent by the user
// Alternatively, you can use input_message_content to send a message with the specified content instead of the voice message.
// Note: This will only work in Telegram versions released after 9 April, 2016
// Older clients will ignore them.
type InlineQueryResultCachedVoice struct {
	Caption string `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Id string `json:"id"`
	InputMessageContent interface{} `json:"input_message_content,omitempty"`
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
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
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
