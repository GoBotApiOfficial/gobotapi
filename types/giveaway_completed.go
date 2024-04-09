// Code AutoGenerated; DO NOT EDIT.

package types

// GiveawayCompleted Represents a service message about the completion of a giveaway without public winners.
type GiveawayCompleted struct {
	GiveawayMessage     *Message `json:"giveaway_message,omitempty"`
	UnclaimedPrizeCount int      `json:"unclaimed_prize_count,omitempty"`
	WinnerCount         int      `json:"winner_count"`
}