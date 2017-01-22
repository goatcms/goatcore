package totype

import "strconv"

func StringToBool(from string) (bool, error) {
	return strconv.ParseBool(from)
}
