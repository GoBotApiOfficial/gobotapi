// Code AutoGenerated; DO NOT EDIT.

package types

import (
	"encoding/json"
	"fmt"
)

// ReactionCount Represents a reaction added to a message along with the number of times it was added.
type ReactionCount struct {
	TotalCount int `json:"total_count"`
	Type       any `json:"type"`
}

func (entity ReactionCount) MarshalJSON() ([]byte, error) {
	switch entity.Type.(type) {
	case ReactionTypeEmoji, ReactionTypeCustomEmoji:
		break
	default:
		return nil, fmt.Errorf("type: unknown type: %T", entity.Type)
	}
	type x0 ReactionCount
	return json.Marshal((x0)(entity))
}
