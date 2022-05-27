package types

type InputInvoiceMessageContent struct {
	Currency string `json:"currency"`
	Description string `json:"description"`
	IsFlexible bool `json:"is_flexible,omitempty"`
	MaxTipAmount int `json:"max_tip_amount,omitempty"`
	NeedEmail bool `json:"need_email,omitempty"`
	NeedName bool `json:"need_name,omitempty"`
	NeedPhoneNumber bool `json:"need_phone_number,omitempty"`
	NeedShippingAddress bool `json:"need_shipping_address,omitempty"`
	Payload string `json:"payload"`
	PhotoHeight int `json:"photo_height,omitempty"`
	PhotoSize int `json:"photo_size,omitempty"`
	PhotoUrl string `json:"photo_url,omitempty"`
	PhotoWidth int64 `json:"photo_width,omitempty"`
	Prices []LabeledPrice `json:"prices,omitempty"`
	ProviderData string `json:"provider_data,omitempty"`
	ProviderToken string `json:"provider_token"`
	SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`
	SuggestedTipAmounts []int `json:"suggested_tip_amounts,omitempty"`
	Title string `json:"title"`
}
