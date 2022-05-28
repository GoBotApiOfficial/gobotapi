package types

// BotCommandScopeChatMember Represents the scope of bot commands, covering a specific member of a group or supergroup chat.
type BotCommandScopeChatMember struct {
	ChatId int64 `json:"chat_id"`
	Type string `json:"type"`
	UserId int64 `json:"user_id"`
}
