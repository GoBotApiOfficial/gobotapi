package types

type Contact struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number"`
	UserId int64 `json:"user_id,omitempty"`
	Vcard string `json:"vcard,omitempty"`
}
