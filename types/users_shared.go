// Code AutoGenerated; DO NOT EDIT.

package types

// UsersShared This object contains information about the users whose identifiers were shared with the bot using a KeyboardButtonRequestUsers button.
type UsersShared struct {
	RequestID int64        `json:"request_id"`
	Users     []SharedUser `json:"users,omitempty"`
}
