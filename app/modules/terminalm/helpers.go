package terminalm

func fixSpace(s string, l int) string {
	ll := l - len(s)
	prefix := make([]rune, ll)
	for i := range prefix {
		prefix[i] = ' '
	}
	return string(prefix) + s
}
