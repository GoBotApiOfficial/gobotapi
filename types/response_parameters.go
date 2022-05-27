package types

type ResponseParameters struct {
	MigrateToChatId int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter int `json:"retry_after,omitempty"`
}
