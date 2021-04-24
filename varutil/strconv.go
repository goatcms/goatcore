package varutil

import (
	"fmt"
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

// ToMapStringInterface convert map[interface{}]interface{} to map[string]interface{}
func ToMapStringInterface(in map[interface{}]interface{}) (out map[string]interface{}) {
	out = make(map[string]interface{})
	for key, value := range in {
		out[fmt.Sprintf("%v", key)] = value
	}
	return
}

// ToMapInterfaceInterface convert map[string]interface{} to map[interface{}]interface{}
func ToMapInterfaceInterface(in map[string]interface{}) (out map[interface{}]interface{}) {
	out = make(map[interface{}]interface{})
	for key, value := range in {
		out[key] = value
	}
	return
}

// ToStringArr convert []interface{} to []string{}
func ToStringArr(in []interface{}) (out []string) {
	out = make([]string, len(in))
	for key, value := range in {
		out[key] = fmt.Sprintf("%v", value)
	}
	return
}
