package types

type PreCheckoutQuery struct {
	Currency string `json:"currency"`
	From User `json:"from"`
	Id string `json:"id"`
	InvoicePayload string `json:"invoice_payload"`
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
	ShippingOptionId string `json:"shipping_option_id,omitempty"`
	TotalAmount int `json:"total_amount"`
}
