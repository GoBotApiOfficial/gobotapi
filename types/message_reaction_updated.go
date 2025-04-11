// Code AutoGenerated; DO NOT EDIT.

package types

import (
	"encoding/json"
	"fmt"
	"reflect"
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
	nilCheck := func(val any) bool {
		if val == nil {
			return true
		}
		v := reflect.ValueOf(val)
		k := v.Kind()
		switch k {
		case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
			return v.IsNil()
		default:
			return false
		}
	}
	_ = nilCheck
	if nilCheck(entity.User) {
		entity.User = nil
	}
	if nilCheck(entity.ActorChat) {
		entity.ActorChat = nil
	}
	for _, x0 := range entity.OldReaction {
		switch x0.(type) {
		case ReactionTypeEmoji, ReactionTypeCustomEmoji, ReactionTypePaid:
			break
		default:
			return nil, fmt.Errorf("old_reaction: unknown type: %T", x0)
		}
	}
	for _, x0 := range entity.NewReaction {
		switch x0.(type) {
		case ReactionTypeEmoji, ReactionTypeCustomEmoji, ReactionTypePaid:
			break
		default:
			return nil, fmt.Errorf("new_reaction: unknown type: %T", x0)
		}
	}
	type x0 MessageReactionUpdated
	return json.Marshal((x0)(entity))
}
