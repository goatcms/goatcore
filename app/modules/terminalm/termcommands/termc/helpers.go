package termc

func fixSpace(s string, l int) string {
	ll := l - len(s)
	prefix := make([]rune, ll)
	for i := range prefix {
		prefix[i] = ' '
	}
	return string(prefix) + s
}

func maxLength(strs []string) (max int) {
	max = 0
	for _, str := range strs {
		if len(str) > max {
			max = len(str)
		}
	}
	return
}
