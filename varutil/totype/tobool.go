package totype

import "strconv"

// StringToBool convert string to boolean
func StringToBool(from string) (bool, error) {
	return strconv.ParseBool(from)
}
