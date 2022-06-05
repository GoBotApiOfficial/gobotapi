package filters

import (
	"github.com/Squirrel-Network/gobotapi/filters"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
)

// IsAdmin Returns a filter that checks if the user is an admin
func (ctx *Wrapper) IsAdmin() filters.FilterOperand {
	return func(options ...any) bool {
		var chatID int64
		var userID int64
		for _, option := range options {
			if chat, ok := option.(*types.Chat); ok && chat != nil {
				chatID = chat.ID
			}
			if user, ok := option.(*types.User); ok && user != nil {
				userID = user.ID
			}

			// For callback queries (grab the chat id from the linked message)
			if message, ok := option.(*types.Message); ok && message != nil {
				chatID = message.Chat.ID
			}
		}
		if chatID != 0 && userID != 0 {
			if ctx.listUsers[chatID] == nil || ctx.listUsers[chatID][userID] == nil {
				invoke, err := ctx.client.Invoke(&methods.GetChatAdministrators{
					ChatID: chatID,
				})
				if err == nil {
					test := invoke.Result.([]types.ChatMember)
					for _, member := range test {
						if ctx.listUsers[chatID] == nil {
							ctx.listUsers[chatID] = make(map[int64]*types.ChatMember)
						}
						ctx.listUsers[chatID][member.User.ID] = &member
					}
				} else {
					return false
				}
			}
			statusMember := ctx.listUsers[chatID][userID].Kind()
			return statusMember == types.TypeChatMemberOwner ||
				statusMember == types.TypeChatMemberAdministrator
		}
		return false
	}
}
