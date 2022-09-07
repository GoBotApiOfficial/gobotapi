package filters

func Or(options ...FilterOperand) FilterOperand {
	return func(values *DataFilter) bool {
		for _, option := range options {
			if option(values) {
				return true
			}
		}
		return false
	}
}
