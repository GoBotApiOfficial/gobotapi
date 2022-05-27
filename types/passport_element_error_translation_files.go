package types

type PassportElementErrorTranslationFiles struct {
	FileHashes []string `json:"file_hashes,omitempty"`
	Message string `json:"message"`
	Source string `json:"source"`
	Type string `json:"type"`
}
