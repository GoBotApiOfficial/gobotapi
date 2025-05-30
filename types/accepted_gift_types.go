// Code AutoGenerated; DO NOT EDIT.

package types

// AcceptedGiftTypes This object describes the types of gifts that can be gifted to a user or a chat.
type AcceptedGiftTypes struct {
	LimitedGifts        bool `json:"limited_gifts"`
	PremiumSubscription bool `json:"premium_subscription"`
	UniqueGifts         bool `json:"unique_gifts"`
	UnlimitedGifts      bool `json:"unlimited_gifts"`
}
