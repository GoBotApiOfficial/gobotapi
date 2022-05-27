package types

type GameHighScore struct {
	Position int `json:"position"`
	Score int `json:"score"`
	User User `json:"user"`
}
