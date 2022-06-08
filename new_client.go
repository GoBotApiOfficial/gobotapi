package gobotapi

func NewClient(token string) *Client {
	return &Client{
		Token: token,
	}
}
