package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

type GetUserProfilePhotos struct {
	Limit int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
	UserId int64 `json:"user_id"`
}

func (entity *GetUserProfilePhotos) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (GetUserProfilePhotos) MethodName() string {
	return "getUserProfilePhotos"
}

func (GetUserProfilePhotos) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.UserProfilePhotos `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeUserProfilePhotos,
		Result: x1.Result,
	}
	return &result, nil
}
