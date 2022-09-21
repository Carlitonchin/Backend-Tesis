package some_utils

func AtLeastOneError(errors ...error) bool {
	for _, err := range errors {
		if err != nil {
			return true
		}
	}

	return false
}
