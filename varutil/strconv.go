package varutil

import (
	"strconv"
	"strings"
)

// Quote transform string pointer to quoter string or "null" const (for nil pointer).
func Quote(s *string) string {
	if s == nil {
		return "null"
	}
	return strconv.Quote(*s)
}

// QuoteArray transform string array to quoter strings
func QuoteArray(s []string, sep string) string {
	quoted := make([]string, len(s))
	for i, value := range s {
		quoted[i] = strconv.Quote(value)
	}
	return strings.Join(quoted, sep)
}

// FormatInt convert int to string if value is set. Otherwise return "null" string
func FormatInt(i *int64, base int) string {
	if i == nil {
		return "null"
	}
	return strconv.FormatInt(*i, base)
}

// FormatIntArray transform int array to formatted string
func FormatIntArray(arr []int64, base int, sep string) string {
	quoted := make([]string, len(arr))
	for i, value := range arr {
		quoted[i] = strconv.FormatInt(value, base)
	}
	return strings.Join(quoted, sep)
}
