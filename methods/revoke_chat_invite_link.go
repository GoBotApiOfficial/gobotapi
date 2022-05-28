package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

// RevokeChatInviteLink Use this method to revoke an invite link created by the bot
// If the primary link is revoked, a new link is automatically generated
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights
// Returns the revoked invite link as ChatInviteLink object.
type RevokeChatInviteLink struct {
	ChatId int64 `json:"chat_id"`
	InviteLink string `json:"invite_link"`
}

func (entity *RevokeChatInviteLink) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
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
	result := rawTypes.Result {
		Kind: types.TypeChatInviteLink,
		Result: x1.Result,
	}
	return &result, nil
}
