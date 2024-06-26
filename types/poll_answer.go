// Code AutoGenerated; DO NOT EDIT.

package types

// PollAnswer Represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	OptionIDs []int64 `json:"option_ids,omitempty"`
	PollID    string  `json:"poll_id"`
	User      *User   `json:"user,omitempty"`
	VoterChat *Chat   `json:"voter_chat,omitempty"`
}
