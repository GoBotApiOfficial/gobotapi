package types

type Poll struct {
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`
	CloseDate int64 `json:"close_date,omitempty"`
	CorrectOptionId int64 `json:"correct_option_id,omitempty"`
	Explanation string `json:"explanation,omitempty"`
	ExplanationEntities []MessageEntity `json:"explanation_entities,omitempty"`
	Id string `json:"id"`
	IsAnonymous bool `json:"is_anonymous"`
	IsClosed bool `json:"is_closed"`
	OpenPeriod int `json:"open_period,omitempty"`
	Options []PollOption `json:"options,omitempty"`
	Question string `json:"question"`
	TotalVoterCount int `json:"total_voter_count"`
	Type string `json:"type"`
}
