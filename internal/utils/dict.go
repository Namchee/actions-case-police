package utils

// MergeDictionary fills a map with entries from other map.
// if a key from map `src` already exists in map `dest`, the value
// for that key will be replaced
func MergeDictionary(dest *map[string]string, src *map[string]string) {
	for k, v := range *src {
		(*dest)[k] = v
	}
}

// RemoveEntries deletes all keys found in `exclusion` from `src`
func RemoveEntries(src *map[string]string, exclusion []string) {
	for _, val := range exclusion {
		delete(*src, val)
	}
}
