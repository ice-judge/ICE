package utils

import "os"

// Getenv returns the enviornment variable named by the key. If the enviornment
// variable was not set, this will return default_.
func Getenv(key string, default_ string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		return default_
	}
	return val
}
