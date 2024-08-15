// Code AutoGenerated; DO NOT EDIT.

package types

import (
	"encoding/json"
	"fmt"
)

// InlineQueryResultVoice Represents a link to a voice recording in an .OGG container encoded with OPUS
// By default, this voice recording will be sent by the user
// Alternatively, you can use input_message_content to send a message with the specified content instead of the the voice message.
type InlineQueryResultVoice struct {
	Caption             string                `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ID                  string                `json:"id"`
	InputMessageContent any                   `json:"input_message_content,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Title               string                `json:"title"`
	VoiceDuration       int                   `json:"voice_duration,omitempty"`
	VoiceURL            string                `json:"voice_url"`
}

func (entity InlineQueryResultVoice) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type                string                `json:"type"`
		ID                  string                `json:"id"`
		VoiceURL            string                `json:"voice_url"`
		Title               string                `json:"title"`
		Caption             string                `json:"caption,omitempty"`
		ParseMode           string                `json:"parse_mode,omitempty"`
		CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
		VoiceDuration       int                   `json:"voice_duration,omitempty"`
		ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent any                   `json:"input_message_content,omitempty"`
	}{
		Type:                "voice",
		ID:                  entity.ID,
		VoiceURL:            entity.VoiceURL,
		Title:               entity.Title,
		Caption:             entity.Caption,
		ParseMode:           entity.ParseMode,
		CaptionEntities:     entity.CaptionEntities,
		VoiceDuration:       entity.VoiceDuration,
		ReplyMarkup:         entity.ReplyMarkup,
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
