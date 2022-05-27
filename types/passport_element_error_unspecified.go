package types

type PassportElementErrorUnspecified struct {
	ElementHash string `json:"element_hash"`
	Message string `json:"message"`
	Source string `json:"source"`
	Type string `json:"type"`
}
