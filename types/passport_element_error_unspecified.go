package types

// PassportElementErrorUnspecified Represents an issue in an unspecified place
// The error is considered resolved when new data is added.
type PassportElementErrorUnspecified struct {
	ElementHash string `json:"element_hash"`
	Message string `json:"message"`
	Source string `json:"source"`
	Type string `json:"type"`
}
