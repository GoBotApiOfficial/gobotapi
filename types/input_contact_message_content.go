package types

type InputContactMessageContent struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number"`
	Vcard string `json:"vcard,omitempty"`
}
