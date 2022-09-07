package filters

func Not(option FilterOperand) FilterOperand {
	return func(values *DataFilter) bool {
		return !option(values)
	}
}
