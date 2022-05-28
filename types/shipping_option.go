package types

// ShippingOption Represents one shipping option.
type ShippingOption struct {
	Id string `json:"id"`
	Prices []LabeledPrice `json:"prices,omitempty"`
	Title string `json:"title"`
}
