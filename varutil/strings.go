package varutil

import (
	"strings"
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
