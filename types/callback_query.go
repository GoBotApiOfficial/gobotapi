package types

type CallbackQuery struct {
	ChatInstance string `json:"chat_instance"`
	Data string `json:"data,omitempty"`
	From User `json:"from"`
	GameShortName string `json:"game_short_name,omitempty"`
	Id string `json:"id"`
	InlineMessageId string `json:"inline_message_id,omitempty"`
	Message *Message `json:"message,omitempty"`
}
