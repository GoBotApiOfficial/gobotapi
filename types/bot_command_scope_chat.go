package types

// BotCommandScopeChat Represents the scope of bot commands, covering a specific chat.
type BotCommandScopeChat struct {
	ChatId int64 `json:"chat_id"`
	Type string `json:"type"`
}
