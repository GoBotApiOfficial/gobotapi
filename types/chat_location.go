package types

// ChatLocation Represents a location to which a chat is connected.
type ChatLocation struct {
	Address string `json:"address"`
	Location Location `json:"location"`
}
