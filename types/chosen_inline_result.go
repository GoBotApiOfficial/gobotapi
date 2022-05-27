package types

type ChosenInlineResult struct {
	From User `json:"from"`
	InlineMessageId string `json:"inline_message_id,omitempty"`
	Location *Location `json:"location,omitempty"`
	Query string `json:"query"`
	ResultId string `json:"result_id"`
}
