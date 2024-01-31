// Code AutoGenerated; DO NOT EDIT.

package types

// ChatBoostRemoved Represents a boost removed from a chat.
type ChatBoostRemoved struct {
	BoostID    string          `json:"boost_id"`
	Chat       Chat            `json:"chat"`
	RemoveDate int64           `json:"remove_date"`
	Source     ChatBoostSource `json:"source"`
}
