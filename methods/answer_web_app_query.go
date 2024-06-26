// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// AnswerWebAppQuery Use this method to set the result of an interaction with a Web App and send a corresponding message on behalf of the user to the chat from which the query originated
// On success, a SentWebAppMessage object is returned.
type AnswerWebAppQuery struct {
	Result        types.InlineQueryResult `json:"result"`
	WebAppQueryID string                  `json:"web_app_query_id"`
}

func (entity *AnswerWebAppQuery) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *AnswerWebAppQuery) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (AnswerWebAppQuery) MethodName() string {
	return "answerWebAppQuery"
}

func (AnswerWebAppQuery) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.SentWebAppMessage `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeSentWebAppMessage,
		Result: x1.Result,
	}
	return &result, nil
}
