package types

type PassportData struct {
	Credentials EncryptedCredentials `json:"credentials"`
	Data []EncryptedPassportElement `json:"data,omitempty"`
}
