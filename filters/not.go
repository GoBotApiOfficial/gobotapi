package filters

func Not(option FilterOperand) FilterOperand {
	return func(values ...any) bool {
		return !option(values)
	}
}
