package varg

import (
	"strings"

	"github.com/goatcms/goatcore/varutil/goaterr"
)

// MatchString return matched value. Change value to lower. Or error if value is incorrect
func MatchString(name, value string, allow []string, defaultValue string) (result string, err error) {
	value = strings.ToLower(value)
	if value == "" {
		return defaultValue, nil
	}
	for _, row := range allow {
		if strings.ToLower(row) == value {
			return value, nil
		}
	}
	return "", goaterr.Errorf("Incorrect value %v for %s. Allow values: '%v'", value, name, strings.Join(allow, "', '"))
}
