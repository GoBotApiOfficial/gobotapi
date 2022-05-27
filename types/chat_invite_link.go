package types

type ChatInviteLink struct {
	CreatesJoinRequest bool `json:"creates_join_request"`
	Creator User `json:"creator"`
	ExpireDate int64 `json:"expire_date,omitempty"`
	InviteLink string `json:"invite_link"`
	IsPrimary bool `json:"is_primary"`
	IsRevoked bool `json:"is_revoked"`
	MemberLimit int `json:"member_limit,omitempty"`
	Name string `json:"name,omitempty"`
	PendingJoinRequestCount int `json:"pending_join_request_count,omitempty"`
}
