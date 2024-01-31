package filters

import "github.com/GoBotApiOfficial/gobotapi"

func Filter[T Filterable](callable func(*gobotapi.Client, T), operands FilterOperand) func(*gobotapi.Client, T) {
	return func(client *gobotapi.Client, update T) {
		if operands(filterableData(client, update)) {
			callable(client, update)
		}
	}
}
