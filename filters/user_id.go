package filters

import "github.com/Squirrel-Network/gobotapi/types"

func UserID(id int64) FilterOperand {
	return func(values ...any) bool {
		for _, value := range values {
			if user, ok := value.(*types.User); ok {
				return user.ID == id
			}
		}
		return false
	}
}
