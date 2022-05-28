package types

// PollAnswer Represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	OptionIds []int64 `json:"option_ids,omitempty"`
	PollId string `json:"poll_id"`
	User User `json:"user"`
}
