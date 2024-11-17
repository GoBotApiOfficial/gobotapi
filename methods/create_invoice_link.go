// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// CreateInvoiceLink Use this method to create a link for an invoice
// Returns the created invoice link as String on success.
type CreateInvoiceLink struct {
	BusinessConnectionID      string               `json:"business_connection_id,omitempty"`
	Currency                  string               `json:"currency"`
	Description               string               `json:"description"`
	IsFlexible                bool                 `json:"is_flexible,omitempty"`
	MaxTipAmount              int                  `json:"max_tip_amount,omitempty"`
	NeedEmail                 bool                 `json:"need_email,omitempty"`
	NeedName                  bool                 `json:"need_name,omitempty"`
	NeedPhoneNumber           bool                 `json:"need_phone_number,omitempty"`
	NeedShippingAddress       bool                 `json:"need_shipping_address,omitempty"`
	Payload                   string               `json:"payload"`
	PhotoHeight               int                  `json:"photo_height,omitempty"`
	PhotoSize                 int                  `json:"photo_size,omitempty"`
	PhotoURL                  string               `json:"photo_url,omitempty"`
	PhotoWidth                int64                `json:"photo_width,omitempty"`
	Prices                    []types.LabeledPrice `json:"prices,omitempty"`
	ProviderData              string               `json:"provider_data,omitempty"`
	ProviderToken             string               `json:"provider_token,omitempty"`
	SendEmailToProvider       bool                 `json:"send_email_to_provider,omitempty"`
	SendPhoneNumberToProvider bool                 `json:"send_phone_number_to_provider,omitempty"`
	SubscriptionPeriod        int                  `json:"subscription_period,omitempty"`
	SuggestedTipAmounts       []int                `json:"suggested_tip_amounts,omitempty"`
	Title                     string               `json:"title"`
}

func (entity *CreateInvoiceLink) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *CreateInvoiceLink) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (CreateInvoiceLink) MethodName() string {
	return "createInvoiceLink"
}

func (CreateInvoiceLink) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result string `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeString,
		Result: x1.Result,
	}
	return &result, nil
}
