package types

type LoginUrl struct {
	BotUsername string `json:"bot_username,omitempty"`
	ForwardText string `json:"forward_text,omitempty"`
	RequestWriteAccess bool `json:"request_write_access,omitempty"`
	Url string `json:"url"`
}
