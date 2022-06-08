package filters

import "github.com/Squirrel-Network/gobotapi"

func Filter[T Filterable](callable func(*gobotapi.Client, T), operands FilterOperand) func(*gobotapi.Client, T) {
	return func(client *gobotapi.Client, update T) {
		if operands(filterableData(update)) {
			callable(client, update)
		}
	}
}
