package filters

import (
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/types"
)

type GenericClient interface {
	OnMessage(func(*gobotapi.Client, types.Message))
	OnChatMember(func(*gobotapi.Client, types.ChatMemberUpdated))
}
