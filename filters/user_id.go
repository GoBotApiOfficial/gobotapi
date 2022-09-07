package filters

import "github.com/Squirrel-Network/gobotapi/utils"

func UserID(idList ...int64) FilterOperand {
	return func(values *DataFilter) bool {
		if values.From != nil {
			return utils.Contains(idList, values.From.ID)
		}
		return false
	}
}
