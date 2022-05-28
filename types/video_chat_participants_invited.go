package types

// VideoChatParticipantsInvited Represents a service message about new members invited to a video chat.
type VideoChatParticipantsInvited struct {
	Users []User `json:"users,omitempty"`
}
