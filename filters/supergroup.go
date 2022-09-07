package filters

func SuperGroup() FilterOperand {
	return func(values *DataFilter) bool {
		if values.Chat != nil {
			return values.Chat.Type == "supergroup"
		}
		if values.Message != nil {
			return values.Message.Chat.Type == "supergroup"
		}
		return false
	}
}
