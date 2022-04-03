package utils

// ContainsString checks if `v` is present on slice `s`
func ContainsString(s []string, v string) bool {
	for _, val := range s {
		if val == v {
			return true
		}
	}

	return false
}
