package methods

import "github.com/Squirrel-Network/gobotapi/types"
import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "encoding/json"

type AnswerShippingQuery struct {
	ErrorMessage string `json:"error_message,omitempty"`
	Ok bool `json:"ok"`
	ShippingOptions []types.ShippingOption `json:"shipping_options,omitempty"`
	ShippingQueryId string `json:"shipping_query_id"`
}

func (entity *AnswerShippingQuery) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (AnswerShippingQuery) MethodName() string {
	return "answerShippingQuery"
}

func (AnswerShippingQuery) ParseResult(response []byte) (*rawTypes.Result, error) {
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
