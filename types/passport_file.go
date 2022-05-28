package types

// PassportFile Represents a file uploaded to Telegram Passport
// Currently all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
	FileDate int64 `json:"file_date"`
	FileId string `json:"file_id"`
	FileSize int `json:"file_size"`
	FileUniqueId string `json:"file_unique_id"`
}
