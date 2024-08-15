// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"reflect"
)

// PinChatMessage Use this method to add a message to the list of pinned messages in a chat
// If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel
// Returns True on success.
type PinChatMessage struct {
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	ChatID               any    `json:"chat_id"`
	DisableNotification  bool   `json:"disable_notification,omitempty"`
	MessageID            int64  `json:"message_id"`
}

func (entity *PinChatMessage) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *PinChatMessage) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity PinChatMessage) MarshalJSON() ([]byte, error) {
	if !reflect.DeepEqual(entity.ChatID, nil) {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 PinChatMessage
	return json.Marshal((x0)(entity))
}

func (PinChatMessage) MethodName() string {
	return "pinChatMessage"
}

func (PinChatMessage) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result bool `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeBoolean,
		Result: x1.Result,
	}
	return &result, nil
}
