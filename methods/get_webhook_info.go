package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

type GetWebhookInfo struct{}

func (entity *GetWebhookInfo) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (GetWebhookInfo) MethodName() string {
	return "getWebhookInfo"
}

func (GetWebhookInfo) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.WebhookInfo `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeWebhookInfo,
		Result: x1.Result,
	}
	return &result, nil
}
