package varutil

import "strings"

// GOPath return golang path like host.com/user/repo
// for example: github.com/goatcms/goatcore
// If it is not external path return empty string.
func GOPath(repourl string) (w string) {
	var (
		e     []string
		index int
	)
	if index = strings.Index(repourl, "://"); index != -1 {
		repourl = repourl[index+3:]
	}
	if strings.HasSuffix(repourl, ".git") {
		repourl = repourl[0 : len(repourl)-4]
	}
	if e = strings.Split(repourl, "/"); len(e) < 3 {
		return ""
	}
	return e[0] + "/" + e[1] + "/" + e[2]
}
