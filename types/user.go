package types

type User struct {
	CanJoinGroups bool `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages,omitempty"`
	FirstName string `json:"first_name"`
	Id int64 `json:"id"`
	IsBot bool `json:"is_bot"`
	LanguageCode string `json:"language_code,omitempty"`
	LastName string `json:"last_name,omitempty"`
	SupportsInlineQueries bool `json:"supports_inline_queries,omitempty"`
	Username string `json:"username,omitempty"`
}
