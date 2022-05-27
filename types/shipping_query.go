package types

type ShippingQuery struct {
	From User `json:"from"`
	Id string `json:"id"`
	InvoicePayload string `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}
