package filters

import (
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/types"
)

func FilterWrapper(client GenericClient) *Wrapper {
	wrapper := &Wrapper{
		listUsers: make(map[int64]map[int64]*types.ChatMember),
	}

	// Listen for user was promoted or demoted in a chat
	client.OnChatMember(func(_ *gobotapi.Client, update types.ChatMemberUpdated) {
		if wrapper.listUsers[update.Chat.ID] == nil {
			wrapper.listUsers[update.Chat.ID] = make(map[int64]*types.ChatMember)
		}
		wrapper.listUsers[update.Chat.ID][update.NewChatMember.User.ID] = &update.NewChatMember
	})

	// Listen for users that left the chat
	client.OnMessage(func(_ *gobotapi.Client, update types.Message) {
		if update.LeftChatMember != nil {
			delete(wrapper.listUsers[update.Chat.ID], update.LeftChatMember.ID)
		}
	})
	return wrapper
}
