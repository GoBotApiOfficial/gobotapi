package filters

import (
	"github.com/GoBotApiOfficial/gobotapi/filters"
	"github.com/GoBotApiOfficial/gobotapi/methods"
	"github.com/GoBotApiOfficial/gobotapi/types"
)

// IsAdmin Returns a filter that checks if the user is an admin
func (ctx *Wrapper) IsAdmin() filters.FilterOperand {
	return func(options *filters.DataFilter) bool {
		ctx.mutex.Lock()
		defer ctx.mutex.Unlock()
		var chatID int64
		var userID int64
		if options.Chat != nil {
			chatID = options.Chat.ID
		}
		if options.From != nil {
			userID = options.From.ID
		}

		// For callback queries (grab the chat id from the linked message)
		if options.Message != nil {
			chatID = options.Message.Chat.ID
		}
		if chatID != 0 && userID != 0 {
			if ctx.listUsers[chatID] == nil || ctx.listUsers[chatID][userID] == nil {
				invoke, err := options.Client.Invoke(&methods.GetChatAdministrators{
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
			member := ctx.listUsers[chatID][userID]
			statusMember := -1
			if member != nil {
				statusMember = member.Kind()
			}
			return statusMember == types.TypeChatMemberOwner ||
				statusMember == types.TypeChatMemberAdministrator
		}
		return false
	}
}
