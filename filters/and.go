package filters

func And(options ...FilterOperand) FilterOperand {
	return func(values ...any) bool {
		for _, option := range options {
			if !option(values...) {
				return false
			}
		}
		return true
	}
}
