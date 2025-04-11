// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// EditStory Edits a story previously posted by the bot on behalf of a managed business account
// Requires the can_manage_stories business bot right
// Returns Story on success.
type EditStory struct {
	Areas                []types.StoryArea       `json:"areas,omitempty"`
	BusinessConnectionID string                  `json:"business_connection_id"`
	Caption              string                  `json:"caption,omitempty"`
	CaptionEntities      []types.MessageEntity   `json:"caption_entities,omitempty"`
	Content              types.InputStoryContent `json:"content"`
	ParseMode            string                  `json:"parse_mode,omitempty"`
	StoryID              int64                   `json:"story_id"`
}

func (entity *EditStory) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *EditStory) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (EditStory) MethodName() string {
	return "editStory"
}

func (EditStory) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.Story `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeStory,
		Result: x1.Result,
	}
	return &result, nil
}
