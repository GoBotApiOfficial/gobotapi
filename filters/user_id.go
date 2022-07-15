package filters

import (
	"github.com/Squirrel-Network/gobotapi/types"
	"github.com/Squirrel-Network/gobotapi/utils"
)

func UserID(idList ...int64) FilterOperand {
	return func(values ...any) bool {
		for _, value := range values {
			if user, ok := value.(*types.User); ok && ok != nil {
				return utils.Contains(idList, user.ID)
			}
		}
		return false
	}
}
