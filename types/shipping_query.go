package types

// ShippingQuery This object contains information about an incoming shipping query.
type ShippingQuery struct {
	From User `json:"from"`
	Id string `json:"id"`
	InvoicePayload string `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}
