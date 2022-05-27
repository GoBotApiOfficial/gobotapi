package types

type PollAnswer struct {
	OptionIds []int64 `json:"option_ids,omitempty"`
	PollId string `json:"poll_id"`
	User User `json:"user"`
}
