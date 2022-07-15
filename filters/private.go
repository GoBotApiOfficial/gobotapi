package filters

import "github.com/Squirrel-Network/gobotapi/types"

func Private() FilterOperand {
	return func(values ...any) bool {
		for _, value := range values {
			if chat, ok := value.(*types.Chat); ok && ok != nil {
				return chat.Type == "private"
			}
			if message, ok := value.(*types.Message); ok && ok != nil {
				return message.Chat.Type == "private"
			}
		}
		return false
	}
}
