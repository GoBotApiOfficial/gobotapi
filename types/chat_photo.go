package types

type ChatPhoto struct {
	BigFileId string `json:"big_file_id"`
	BigFileUniqueId string `json:"big_file_unique_id"`
	SmallFileId string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
}
