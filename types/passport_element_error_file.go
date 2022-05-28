package types

// PassportElementErrorFile Represents an issue with a document scan
// The error is considered resolved when the file with the document scan changes.
type PassportElementErrorFile struct {
	FileHash string `json:"file_hash"`
	Message string `json:"message"`
	Source string `json:"source"`
	Type string `json:"type"`
}
