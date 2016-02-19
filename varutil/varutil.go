package varutil

func IsInStringArray(array []string, value string) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			return true
		}
	}
	return false
}
