// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"gobotapi/types"
	rawTypes "gobotapi/types/raw"
)

// RevokeChatInviteLink Use this method to revoke an invite link created by the bot
// If the primary link is revoked, a new link is automatically generated
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights
// Returns the revoked invite link as ChatInviteLink object.
type RevokeChatInviteLink struct {
	ChatID     any    `json:"chat_id"`
	InviteLink string `json:"invite_link"`
}

func (entity *RevokeChatInviteLink) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *RevokeChatInviteLink) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity RevokeChatInviteLink) MarshalJSON() ([]byte, error) {
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 RevokeChatInviteLink
	return json.Marshal((x0)(entity))
}

func (RevokeChatInviteLink) MethodName() string {
	return "revokeChatInviteLink"
}

func (RevokeChatInviteLink) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.ChatInviteLink `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeChatInviteLink,
		Result: x1.Result,
	}
	return &result, nil
}
