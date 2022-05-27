package types

type PassportElementErrorTranslationFile struct {
	FileHash string `json:"file_hash"`
	Message string `json:"message"`
	Source string `json:"source"`
	Type string `json:"type"`
}
