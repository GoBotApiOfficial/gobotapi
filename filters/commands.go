package filters

func Commands(command []string, aliasList ...string) FilterOperand {
	return func(values *DataFilter) bool {
		for _, cmd := range command {
			if Command(cmd, aliasList...)(values) {
				return true
			}
		}
		return false
	}
}
