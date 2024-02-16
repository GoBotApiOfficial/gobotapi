// Code AutoGenerated; DO NOT EDIT.

package types

import (
	"encoding/json"
	"fmt"
)

// MessageReactionUpdated Represents a change of a reaction on a message performed by a user.
type MessageReactionUpdated struct {
	ActorChat   *Chat          `json:"actor_chat,omitempty"`
	Chat        Chat           `json:"chat"`
	Date        int64          `json:"date"`
	MessageID   int64          `json:"message_id"`
	NewReaction []ReactionType `json:"new_reaction,omitempty"`
	OldReaction []ReactionType `json:"old_reaction,omitempty"`
	User        *User          `json:"user,omitempty"`
}

func (entity MessageReactionUpdated) MarshalJSON() ([]byte, error) {
	for _, x0 := range entity.OldReaction {
		switch x0.(type) {
		case ReactionTypeEmoji, ReactionTypeCustomEmoji:
			break
		default:
			return nil, fmt.Errorf("old_reaction: unknown type: %T", x0)
		}
	}
	for _, x0 := range entity.NewReaction {
		switch x0.(type) {
		case ReactionTypeEmoji, ReactionTypeCustomEmoji:
			break
		default:
			return nil, fmt.Errorf("new_reaction: unknown type: %T", x0)
		}
	}
	type x0 MessageReactionUpdated
	return json.Marshal((x0)(entity))
}
