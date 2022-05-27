package types

type PassportElementErrorDataField struct {
	DataHash string `json:"data_hash"`
	FieldName string `json:"field_name"`
	Message string `json:"message"`
	Source string `json:"source"`
	Type string `json:"type"`
}
