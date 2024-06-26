package filters

import "github.com/GoBotApiOfficial/gobotapi/utils"

func ChatID(idList ...int64) FilterOperand {
	return func(values *DataFilter) bool {
		if values.Chat != nil {
			return utils.Contains(idList, values.Chat.ID)
		}
		if values.Message != nil {
			return utils.Contains(idList, values.Message.Chat.ID)
		}
		return false
	}
}
