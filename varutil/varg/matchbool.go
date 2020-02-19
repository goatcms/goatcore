package varg

// MatchBool return matched value. Change value to lower. Or error if value is incorrect
func MatchBool(name, value string, defaultValue bool) (result bool, err error) {
	if value == "" {
		return defaultValue, nil
	}
	if value, err = MatchString(name, value, []string{
		"true",
		"false",
	}, ""); err != nil {
		return false, err
	}
	return value == "true", nil
}
