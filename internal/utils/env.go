package utils

import (
	"os"
	"strconv"
	"strings"
)

// ReadEnvBool read and parse boolean environment variables.
// Will return `true` if the variable cannot be parsed
func ReadEnvBool(key string) bool {
	value := os.Getenv(key)
	parsed, err := strconv.ParseBool(value)

	if err != nil {
		return true
	}

	return parsed
}

// ReadEnvString read string environment variables.
// This function is a wrapper function to 'os.Getenv()'
func ReadEnvString(key string) string {
	return os.Getenv(key)
}

// ReadEnvStringArray read an array of string environment variables.
// Must be comma-separated
// Automatically trims all values
func ReadEnvStringArray(key string) []string {
	raw := os.Getenv(key)

	if len(raw) == 0 {
		return []string{}
	}

	tokens := strings.Split(raw, ",")

	var value []string

	for i := range tokens {
		value = append(value, strings.TrimSpace(tokens[i]))
	}

	return value
}
