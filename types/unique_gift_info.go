// Code AutoGenerated; DO NOT EDIT.

package types

// UniqueGiftInfo Describes a service message about a unique gift that was sent or received.
type UniqueGiftInfo struct {
	Gift              UniqueGift `json:"gift"`
	Origin            string     `json:"origin"`
	OwnedGiftID       string     `json:"owned_gift_id,omitempty"`
	TransferStarCount int        `json:"transfer_star_count,omitempty"`
}
