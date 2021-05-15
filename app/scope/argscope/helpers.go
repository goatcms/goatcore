package argscope

// SeparateArgs return separated args and isolated arguments after '--' separate sequnce
func SeparateArgs(all []string) (args, separated []string) {
	args = all
	index := -1
	for i, val := range all {
		if val == "--" {
			index = i
			break
		}
	}
	if index != -1 && len(all) > index {
		separated = all[index+1:]
	}
	if index != -1 {
		args = args[:index]
	}
	return
}
