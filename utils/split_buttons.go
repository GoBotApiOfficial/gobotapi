package utils

import "github.com/GoBotApiOfficial/gobotapi/types"

func SplitButtons(splitBy int, buttons ...types.InlineKeyboardButton) [][]types.InlineKeyboardButton {
	var result [][]types.InlineKeyboardButton
	var tmp []types.InlineKeyboardButton
	for _, button := range buttons {
		tmp = append(tmp, button)
		if len(tmp) == splitBy {
			result = append(result, tmp)
			tmp = nil
		}
	}
	if len(tmp) > 0 {
		result = append(result, tmp)
	}
	return result
}
