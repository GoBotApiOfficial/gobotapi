package types

type ChatMemberUpdated struct {
	Chat Chat `json:"chat"`
	Date int64 `json:"date"`
	From User `json:"from"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
	NewChatMember ChatMember `json:"new_chat_member"`
	OldChatMember ChatMember `json:"old_chat_member"`
}
