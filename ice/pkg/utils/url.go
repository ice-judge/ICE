package utils

import "strings"

// IsHTTPURL returns true if url starts with
// http:// or https://
func IsHTTPURL(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}
