package types

// WebAppData Contains data sent from a Web App to the bot.
type WebAppData struct {
	ButtonText string `json:"button_text"`
	Data string `json:"data"`
}
