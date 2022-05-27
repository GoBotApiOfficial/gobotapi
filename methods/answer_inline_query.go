package methods

import "github.com/Squirrel-Network/gobotapi/types"
import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "encoding/json"

type AnswerInlineQuery struct {
	CacheTime int `json:"cache_time,omitempty"`
	InlineQueryId string `json:"inline_query_id"`
	IsPersonal bool `json:"is_personal,omitempty"`
	NextOffset string `json:"next_offset,omitempty"`
	Results []types.InlineQueryResult `json:"results,omitempty"`
	SwitchPmParameter string `json:"switch_pm_parameter,omitempty"`
	SwitchPmText string `json:"switch_pm_text,omitempty"`
}

func (entity *AnswerInlineQuery) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (AnswerInlineQuery) MethodName() string {
	return "answerInlineQuery"
}

func (AnswerInlineQuery) ParseResult(response []byte) (*rawTypes.Result, error) {
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
