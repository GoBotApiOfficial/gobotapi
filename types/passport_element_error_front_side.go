package types

// PassportElementErrorFrontSide Represents an issue with the front side of a document
// The error is considered resolved when the file with the front side of the document changes.
type PassportElementErrorFrontSide struct {
	FileHash string `json:"file_hash"`
	Message string `json:"message"`
	Source string `json:"source"`
	Type string `json:"type"`
}
