package filters

import (
	"github.com/Squirrel-Network/gobotapi/types"
	"github.com/Squirrel-Network/gobotapi/utils"
)

func UserIDs(ids []int64) FilterOperand {
	return func(values ...any) bool {
		for _, value := range values {
			if user, ok := value.(*types.User); ok {
				return utils.Contains(ids, user.ID)
			}
		}
		return false
	}
}
