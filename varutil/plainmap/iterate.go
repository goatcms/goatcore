package plainmap

import (
	"regexp"
	"sort"
	"strings"
)

// Keys return sub keys for keyPath
func Keys(plainmap map[string]string, keyPath string) (keys []string) {
	m := map[string]bool{}
	for k := range plainmap {
		if strings.HasPrefix(k, keyPath) {
			k = k[len(keyPath):]
			i := strings.Index(k, ".")
			if i == -1 {
				i = len(k)
			}
			k = k[:i]
			m[k] = true
		}
	}
	keys = make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}

// Strain return map filtred by regexp
func Strain(data map[string]string, reg *regexp.Regexp) (result map[string]string, err error) {
	result = make(map[string]string)
	for key, value := range data {
		if reg.MatchString(key) {
			result[key] = value
		}
	}
	return result, nil
}
