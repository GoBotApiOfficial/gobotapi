package types

import "encoding/json"
import "fmt"

// InlineQueryResultArticle Represents a link to an article or web page.
type InlineQueryResultArticle struct {
	Description string `json:"description,omitempty"`
	HideUrl bool `json:"hide_url,omitempty"`
	Id string `json:"id"`
	InputMessageContent interface{} `json:"input_message_content"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbHeight int `json:"thumb_height,omitempty"`
	ThumbUrl string `json:"thumb_url,omitempty"`
	ThumbWidth int64 `json:"thumb_width,omitempty"`
	Title string `json:"title"`
	Url string `json:"url,omitempty"`
}

func (entity InlineQueryResultArticle) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type string `json:"type"`
		Id string `json:"id"`
		Title string `json:"title"`
		InputMessageContent interface{} `json:"input_message_content"`
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		Url string `json:"url,omitempty"`
		HideUrl bool `json:"hide_url,omitempty"`
		Description string `json:"description,omitempty"`
		ThumbUrl string `json:"thumb_url,omitempty"`
		ThumbWidth int64 `json:"thumb_width,omitempty"`
		ThumbHeight int `json:"thumb_height,omitempty"`
	} {
		Type: "article",
		Id: entity.Id,
		Title: entity.Title,
		InputMessageContent: entity.InputMessageContent,
		ReplyMarkup: entity.ReplyMarkup,
		Url: entity.Url,
		HideUrl: entity.HideUrl,
		Description: entity.Description,
		ThumbUrl: entity.ThumbUrl,
		ThumbWidth: entity.ThumbWidth,
		ThumbHeight: entity.ThumbHeight,
	}
	switch entity.InputMessageContent.(type) {
		case InputTextMessageContent, InputLocationMessageContent, InputVenueMessageContent, InputContactMessageContent, InputInvoiceMessageContent:
			break
		default:
			return nil, fmt.Errorf("input_message_content: unknown type: %T", entity.InputMessageContent)
	}
	return json.Marshal(alias)
}
