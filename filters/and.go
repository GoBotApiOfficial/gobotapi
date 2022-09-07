package filters

func And(options ...FilterOperand) FilterOperand {
	return func(values *DataFilter) bool {
		for _, option := range options {
			if !option(values) {
				return false
			}
		}
		return true
	}
}
