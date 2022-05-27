package types

type EncryptedPassportElement struct {
	Data string `json:"data,omitempty"`
	Email string `json:"email,omitempty"`
	Files []PassportFile `json:"files,omitempty"`
	FrontSide *PassportFile `json:"front_side,omitempty"`
	Hash string `json:"hash"`
	PhoneNumber string `json:"phone_number,omitempty"`
	ReverseSide *PassportFile `json:"reverse_side,omitempty"`
	Selfie *PassportFile `json:"selfie,omitempty"`
	Translation []PassportFile `json:"translation,omitempty"`
	Type string `json:"type"`
}
