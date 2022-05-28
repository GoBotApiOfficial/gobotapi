package types

// Game Represents a game
// Use BotFather to create and edit games, their short names will act as unique identifiers.
type Game struct {
	Animation *Animation `json:"animation,omitempty"`
	Description string `json:"description"`
	Photo []PhotoSize `json:"photo,omitempty"`
	Text string `json:"text,omitempty"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	Title string `json:"title"`
}
