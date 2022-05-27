package types

type InlineQuery struct {
	ChatType string `json:"chat_type,omitempty"`
	From User `json:"from"`
	Id string `json:"id"`
	Location *Location `json:"location,omitempty"`
	Offset string `json:"offset"`
	Query string `json:"query"`
}
