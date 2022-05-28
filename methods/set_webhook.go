package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

// SetWebhook Use this method to specify a url and receive incoming updates via an outgoing webhook
// Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url, containing a JSON-serialized Update
// In case of an unsuccessful request, we will give up after a reasonable amount of attempts
// Returns True on success.
// If you'd like to make sure that the Webhook request comes from Telegram, we recommend using a secret path in the URL, e.g
// https://www.example.com/<token>
// Since nobody else knows your bot's token, you can be pretty sure it's us.
type SetWebhook struct {
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
	Certificate rawTypes.InputFile `json:"certificate,omitempty"`
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
	IpAddress string `json:"ip_address,omitempty"`
	MaxConnections int `json:"max_connections,omitempty"`
	Url string `json:"url"`
}

func (entity *SetWebhook) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Certificate.(type) {
		case types.InputFile:
			files["certificate"] = entity.Certificate
			entity.Certificate = nil
	}
	return files
}

func (SetWebhook) MethodName() string {
	return "setWebhook"
}

func (SetWebhook) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result bool `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeBoolean,
		Result: x1.Result,
	}
	return &result, nil
}
