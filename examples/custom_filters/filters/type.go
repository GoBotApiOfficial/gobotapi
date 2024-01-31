package filters

import (
	"github.com/GoBotApiOfficial/gobotapi/types"
	"sync"
)

type Wrapper struct {
	mutex     sync.RWMutex
	listUsers map[int64]map[int64]*types.ChatMember
}
