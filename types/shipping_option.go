package types

type ShippingOption struct {
	Id string `json:"id"`
	Prices []LabeledPrice `json:"prices,omitempty"`
	Title string `json:"title"`
}
