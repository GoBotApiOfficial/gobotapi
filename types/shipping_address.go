package types

// ShippingAddress Represents a shipping address.
type ShippingAddress struct {
	City string `json:"city"`
	CountryCode string `json:"country_code"`
	PostCode string `json:"post_code"`
	State string `json:"state"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
}
