package types

// PassportElementErrorDataField Represents an issue in one of the data fields that was provided by the user
// The error is considered resolved when the field's value changes.
type PassportElementErrorDataField struct {
	DataHash string `json:"data_hash"`
	FieldName string `json:"field_name"`
	Message string `json:"message"`
	Source string `json:"source"`
	Type string `json:"type"`
}
