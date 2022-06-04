package filters

import (
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/types"
)

type Wrapper struct {
	client    gobotapi.Client
	listUsers map[int64]map[int64]*types.ChatMember
}
