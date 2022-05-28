package types

// Dice Represents an animated emoji that displays a random value.
type Dice struct {
	Emoji string `json:"emoji"`
	Value int `json:"value"`
}
