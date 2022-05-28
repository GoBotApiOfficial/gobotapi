package types

// GameHighScore Represents one row of the high scores table for a game.
type GameHighScore struct {
	Position int `json:"position"`
	Score int `json:"score"`
	User User `json:"user"`
}
