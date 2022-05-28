package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

// GetGameHighScores Use this method to get data for high score tables
// Will return the score of the specified user and several of their neighbors in a game
// On success, returns an Array of GameHighScore objects.
type GetGameHighScores struct {
	ChatId int64 `json:"chat_id,omitempty"`
	InlineMessageId string `json:"inline_message_id,omitempty"`
	MessageId int64 `json:"message_id,omitempty"`
	UserId int64 `json:"user_id"`
}

func (entity *GetGameHighScores) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (GetGameHighScores) MethodName() string {
	return "getGameHighScores"
}

func (GetGameHighScores) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result []types.GameHighScore `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeArrayOfGameHighScore,
		Result: x1.Result,
	}
	return &result, nil
}
