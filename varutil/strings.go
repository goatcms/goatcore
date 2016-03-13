package varutil

import (
	"fmt"
	"strings"
	"regexp"
)

func HasOneSuffix(str string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(str, suffix) {
			return true
		}
	}
	return false
}

func IsArrContainStr(arr []string, s string) bool {
	for _, record := range arr {
		if record == s {
			return true
		}
	}
	return false
}

func FixDirPath(path *string) {
	if *path == "" {
		return
	}
	if !strings.HasSuffix(*path, "/") {
		*path = *path + "/"
	}
}

func FixUrl(url *string) error {
	if *url == "" {
		return fmt.Errorf("Incorrect url '", url, "'")
	}
	if strings.HasPrefix(*url, "http://") || strings.HasPrefix(*url, "https://") {
		return nil
	}
	*url = "http://" + *url
	return nil
}

func SplitWhite(s string) ([]string, error) {
	reg, err := regexp.Compile("[ \t]+")
	if err != nil {
		return nil, err
	}
	s = reg.ReplaceAllString(s, " ")
	return strings.Split(s, " "), nil
}
