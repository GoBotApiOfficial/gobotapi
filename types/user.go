// Code AutoGenerated; DO NOT EDIT.

package types

// User Represents a Telegram user or bot.
type User struct {
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu,omitempty"`
	CanConnectToBusiness    bool   `json:"can_connect_to_business,omitempty"`
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
	FirstName               string `json:"first_name"`
	HasMainWebApp           bool   `json:"has_main_web_app,omitempty"`
	ID                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	IsPremium               bool   `json:"is_premium,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	LastName                string `json:"last_name,omitempty"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`
	Username                string `json:"username,omitempty"`
}
