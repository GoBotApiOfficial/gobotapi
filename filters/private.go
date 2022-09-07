package filters

func Private() FilterOperand {
	return func(values *DataFilter) bool {
		if values.Chat != nil {
			return values.Chat.Type == "private"
		}
		if values.Message != nil {
			return values.Message.Chat.Type == "private"
		}
		return false
	}
}
