package filters

func Or(options ...FilterOperand) FilterOperand {
	return func(values ...any) bool {
		for _, option := range options {
			if option(values...) {
				return true
			}
		}
		return false
	}
}
