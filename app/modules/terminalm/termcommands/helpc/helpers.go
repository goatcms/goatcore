package helpc

func maxLength(strs []string) (max int) {
	max = 0
	for _, str := range strs {
		if len(str) > max {
			max = len(str)
		}
	}
	return
}
