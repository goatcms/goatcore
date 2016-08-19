package varutil

import (
	"fmt"
	"regexp"
	"strings"
)

// HasOneSuffix checks any element of array has suffix
func HasOneSuffix(str string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(str, suffix) {
			return true
		}
	}
	return false
}

// IsArrContainStr checks an array contains an element
func IsArrContainStr(arr []string, s string) bool {
	for _, record := range arr {
		if record == s {
			return true
		}
	}
	return false
}

// FixDirPath autocorrect dir path to contains / at its end
func FixDirPath(path *string) {
	if *path == "" {
		return
	}
	if !strings.HasSuffix(*path, "/") {
		*path = *path + "/"
	}
}

// FixURL fix url to start with http://
func FixURL(url *string) error {
	if *url == "" {
		return fmt.Errorf("Incorrect url '%v'", url)
	}
	if strings.HasPrefix(*url, "http://") || strings.HasPrefix(*url, "https://") {
		return nil
	}
	*url = "http://" + *url
	return nil
}

// SplitWhite remove extra spaces
func SplitWhite(s string) ([]string, error) {
	reg, err := regexp.Compile("[ \t]+")
	if err != nil {
		return nil, err
	}
	s = reg.ReplaceAllString(s, " ")
	return strings.Split(s, " "), nil
}
