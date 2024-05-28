// Code AutoGenerated; DO NOT EDIT.

package types

// KeyboardButton Represents one button of the reply keyboard
// At most one of the optional fields must be used to specify type of the button
// For simple text buttons, String can be used instead of this object to specify the button text.
// Note: request_users and request_chat options will only work in Telegram versions released after 3 February, 2023
// Older clients will display unsupported message.
type KeyboardButton struct {
	RequestChat     *KeyboardButtonRequestChat  `json:"request_chat,omitempty"`
	RequestContact  bool                        `json:"request_contact,omitempty"`
	RequestLocation bool                        `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPollType     `json:"request_poll,omitempty"`
	RequestUsers    *KeyboardButtonRequestUsers `json:"request_users,omitempty"`
	Text            string                      `json:"text"`
	WebApp          *WebAppInfo                 `json:"web_app,omitempty"`
}
