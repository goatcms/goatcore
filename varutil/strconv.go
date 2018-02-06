package varutil

import (
	"strconv"
	"strings"
)

func Quote(s *string) string {
	if s == nil {
		return "null"
	}
	return strconv.Quote(*s)
}

func QuoteArray(s []string, sep string) string {
	quoted := make([]string, len(s))
	for i, value := range s {
		quoted[i] = strconv.Quote(value)
	}
	return strings.Join(quoted, sep)
}

func FormatInt(i *int64, base int) string {
	if i == nil {
		return "null"
	}
	return strconv.FormatInt(*i, base)
}

func FormatIntArray(arr []int64, base int, sep string) string {
	quoted := make([]string, len(arr))
	for i, value := range arr {
		quoted[i] = strconv.FormatInt(value, base)
	}
	return strings.Join(quoted, sep)
}
