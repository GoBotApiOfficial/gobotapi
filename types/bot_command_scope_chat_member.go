package types

type BotCommandScopeChatMember struct {
	ChatId int64 `json:"chat_id"`
	Type string `json:"type"`
	UserId int64 `json:"user_id"`
}
