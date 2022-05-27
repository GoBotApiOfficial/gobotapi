package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

type AnswerPreCheckoutQuery struct {
	ErrorMessage string `json:"error_message,omitempty"`
	Ok bool `json:"ok"`
	PreCheckoutQueryId string `json:"pre_checkout_query_id"`
}

func (entity *AnswerPreCheckoutQuery) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (AnswerPreCheckoutQuery) MethodName() string {
	return "answerPreCheckoutQuery"
}

func (AnswerPreCheckoutQuery) ParseResult(response []byte) (*rawTypes.Result, error) {
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
