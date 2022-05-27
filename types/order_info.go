package types

type OrderInfo struct {
	Email string `json:"email,omitempty"`
	Name string `json:"name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}
