package types

type BotCommandScopeChat struct {
	ChatId int64 `json:"chat_id"`
	Type string `json:"type"`
}
