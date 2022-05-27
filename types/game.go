package types

type Game struct {
	Animation *Animation `json:"animation,omitempty"`
	Description string `json:"description"`
	Photo []PhotoSize `json:"photo,omitempty"`
	Text string `json:"text,omitempty"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	Title string `json:"title"`
}
