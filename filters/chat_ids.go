package filters

import (
	"github.com/Squirrel-Network/gobotapi/types"
	"github.com/Squirrel-Network/gobotapi/utils"
)

func ChatIDs(ids []int64) FilterOperand {
	return func(values ...any) bool {
		for _, value := range values {
			if chat, ok := value.(*types.Chat); ok {
				return utils.Contains(ids, chat.ID)
			}
			if message, ok := value.(*types.Message); ok {
				return utils.Contains(ids, message.Chat.ID)
			}
		}
		return false
	}
}
