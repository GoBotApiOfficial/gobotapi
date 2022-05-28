package types

// PassportElementErrorReverseSide Represents an issue with the reverse side of a document
// The error is considered resolved when the file with reverse side of the document changes.
type PassportElementErrorReverseSide struct {
	FileHash string `json:"file_hash"`
	Message string `json:"message"`
	Source string `json:"source"`
	Type string `json:"type"`
}
