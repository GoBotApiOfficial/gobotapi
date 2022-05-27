package types

type ChatJoinRequest struct {
	Bio string `json:"bio,omitempty"`
	Chat Chat `json:"chat"`
	Date int64 `json:"date"`
	From User `json:"from"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}
