package types

// BotCommandScopeChatAdministrators Represents the scope of bot commands, covering all administrators of a specific group or supergroup chat.
type BotCommandScopeChatAdministrators struct {
	ChatId int64 `json:"chat_id"`
	Type string `json:"type"`
}
