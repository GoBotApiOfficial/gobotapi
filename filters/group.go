package filters

func Group() FilterOperand {
	return func(values *DataFilter) bool {
		if values.Chat != nil {
			return values.Chat.Type == "group"
		}
		if values.Message != nil {
			return values.Message.Chat.Type == "group"
		}
		return false
	}
}
