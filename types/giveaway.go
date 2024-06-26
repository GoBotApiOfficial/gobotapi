// Code AutoGenerated; DO NOT EDIT.

package types

// Giveaway Represents a message about a scheduled giveaway.
type Giveaway struct {
	Chats                         []Chat   `json:"chats,omitempty"`
	CountryCodes                  []string `json:"country_codes,omitempty"`
	HasPublicWinners              bool     `json:"has_public_winners,omitempty"`
	OnlyNewMembers                bool     `json:"only_new_members,omitempty"`
	PremiumSubscriptionMonthCount int      `json:"premium_subscription_month_count,omitempty"`
	PrizeDescription              string   `json:"prize_description,omitempty"`
	WinnerCount                   int      `json:"winner_count"`
	WinnersSelectionDate          int64    `json:"winners_selection_date"`
}
