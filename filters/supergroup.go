package filters

import "github.com/Squirrel-Network/gobotapi/types"

func SuperGroup() FilterOperand {
	return func(values ...any) bool {
		for _, value := range values {
			if chat, ok := value.(*types.Chat); ok && chat != nil {
				return chat.Type == "supergroup"
			}
			if message, ok := value.(*types.Message); ok && message != nil {
				return message.Chat.Type == "supergroup"
			}
		}
		return false
	}
}
