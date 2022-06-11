package filters

import (
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

type GenericClient interface {
	OnMessage(func(*gobotapi.Client, types.Message))
	OnChatMember(func(*gobotapi.Client, types.ChatMemberUpdated))
	Invoke(method rawTypes.Method) (*rawTypes.Result, error)
}
