package filters

import "github.com/Squirrel-Network/gobotapi/types"

func ChatID(id int64) FilterOperand {
	return func(values ...any) bool {
		for _, value := range values {
			if chat, ok := value.(*types.Chat); ok {
				return chat.ID == id
			}
			if message, ok := value.(*types.Message); ok {
				return message.Chat.ID == id
			}
		}
		return false
	}
}
