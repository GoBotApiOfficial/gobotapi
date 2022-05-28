package types

// PreCheckoutQuery This object contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	Currency string `json:"currency"`
	From User `json:"from"`
	Id string `json:"id"`
	InvoicePayload string `json:"invoice_payload"`
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
	ShippingOptionId string `json:"shipping_option_id,omitempty"`
	TotalAmount int `json:"total_amount"`
}
