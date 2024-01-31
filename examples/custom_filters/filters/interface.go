package filters

import (
	"github.com/GoBotApiOfficial/gobotapi"
	"github.com/GoBotApiOfficial/gobotapi/types"
)

type GenericClient interface {
	OnMessage(func(*gobotapi.Client, types.Message))
	OnChatMember(func(*gobotapi.Client, types.ChatMemberUpdated))
}
