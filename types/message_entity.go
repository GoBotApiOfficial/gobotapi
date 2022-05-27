package types

type MessageEntity struct {
	Language string `json:"language,omitempty"`
	Length int `json:"length"`
	Offset int `json:"offset"`
	Type string `json:"type"`
	Url string `json:"url,omitempty"`
	User *User `json:"user,omitempty"`
}
