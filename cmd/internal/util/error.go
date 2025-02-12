package util

import "strings"

// FormatErrorToString converts an error to a single-line string by replacing
// newlines and tab characters with spaces.
func FormatErrorToString(err error) string {
	errStr := strings.ReplaceAll(err.Error(), "\n", " ")
	errStr = strings.ReplaceAll(errStr, "\t", " ")
	return errStr
}
