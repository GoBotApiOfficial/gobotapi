package types

type KeyboardButton struct {
	RequestContact bool `json:"request_contact,omitempty"`
	RequestLocation bool `json:"request_location,omitempty"`
	RequestPoll *KeyboardButtonPollType `json:"request_poll,omitempty"`
	Text string `json:"text"`
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}
