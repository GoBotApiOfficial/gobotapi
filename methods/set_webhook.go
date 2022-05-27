package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

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
