// Code AutoGenerated; DO NOT EDIT.

package types

// ForceReply Upon receiving a message with this object, Telegram clients will display a reply interface to the user (act as if the user has selected the bot's message and tapped 'Reply')
// This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode
// Not supported in channels and for messages sent on behalf of a Telegram Business account.
type ForceReply struct {
	ForceReply            bool   `json:"force_reply"`
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	Selective             bool   `json:"selective,omitempty"`
}
