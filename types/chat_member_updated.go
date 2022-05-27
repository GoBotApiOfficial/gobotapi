package types

type ChatMemberUpdated struct {
	Chat Chat `json:"chat"`
	Date int64 `json:"date"`
	From User `json:"from"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
	NewChatMember interface{} `json:"new_chat_member"`
	OldChatMember interface{} `json:"old_chat_member"`
}
