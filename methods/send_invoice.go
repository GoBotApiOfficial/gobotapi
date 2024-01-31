// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// SendInvoice Use this method to send invoices
// On success, the sent Message is returned.
type SendInvoice struct {
	ChatID                    any                         `json:"chat_id"`
	Currency                  string                      `json:"currency"`
	Description               string                      `json:"description"`
	DisableNotification       bool                        `json:"disable_notification,omitempty"`
	IsFlexible                bool                        `json:"is_flexible,omitempty"`
	MaxTipAmount              int                         `json:"max_tip_amount,omitempty"`
	MessageThreadID           int64                       `json:"message_thread_id,omitempty"`
	NeedEmail                 bool                        `json:"need_email,omitempty"`
	NeedName                  bool                        `json:"need_name,omitempty"`
	NeedPhoneNumber           bool                        `json:"need_phone_number,omitempty"`
	NeedShippingAddress       bool                        `json:"need_shipping_address,omitempty"`
	Payload                   string                      `json:"payload"`
	PhotoHeight               int                         `json:"photo_height,omitempty"`
	PhotoSize                 int                         `json:"photo_size,omitempty"`
	PhotoURL                  string                      `json:"photo_url,omitempty"`
	PhotoWidth                int64                       `json:"photo_width,omitempty"`
	Prices                    []types.LabeledPrice        `json:"prices,omitempty"`
	ProtectContent            bool                        `json:"protect_content,omitempty"`
	ProviderData              string                      `json:"provider_data,omitempty"`
	ProviderToken             string                      `json:"provider_token"`
	ReplyMarkup               *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ReplyParameters           *types.ReplyParameters      `json:"reply_parameters,omitempty"`
	SendEmailToProvider       bool                        `json:"send_email_to_provider,omitempty"`
	SendPhoneNumberToProvider bool                        `json:"send_phone_number_to_provider,omitempty"`
	StartParameter            string                      `json:"start_parameter,omitempty"`
	SuggestedTipAmounts       []int                       `json:"suggested_tip_amounts,omitempty"`
	Title                     string                      `json:"title"`
}

func (entity *SendInvoice) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SendInvoice) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity SendInvoice) MarshalJSON() ([]byte, error) {
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 SendInvoice
	return json.Marshal((x0)(entity))
}

func (SendInvoice) MethodName() string {
	return "sendInvoice"
}

func (SendInvoice) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.Message `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeMessage,
		Result: x1.Result,
	}
	return &result, nil
}
