package types

type ChatPermissions struct {
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo bool `json:"can_change_info,omitempty"`
	CanInviteUsers bool `json:"can_invite_users,omitempty"`
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
	CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`
	CanSendMessages bool `json:"can_send_messages,omitempty"`
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`
	CanSendPolls bool `json:"can_send_polls,omitempty"`
}
