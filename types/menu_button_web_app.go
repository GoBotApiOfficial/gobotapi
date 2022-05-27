package types

type MenuButtonWebApp struct {
	Text string `json:"text"`
	Type string `json:"type"`
	WebApp WebAppInfo `json:"web_app"`
}
