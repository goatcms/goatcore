package varutil

import (
	"regexp"
	"strings"

	"github.com/goatcms/goatcore/varutil/goaterr"
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
		return goaterr.Errorf("Incorrect url '%v'", url)
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

// IsWhitespace return true if input is white char
func IsWhitespace(s string) bool {
	return s == " " || s == "\t"
}

// SplitArguments convert string to separated arguments array
func SplitArguments(src string) ([]string, error) {
	var (
		args        []string
		isEscaped   = false
		isSeparated = true
	)
	for i := 0; i < len(src); i++ {
		ch := src[i]
		if ch == ' ' || ch == '\t' {
			isEscaped = false
			isSeparated = true
			continue
		} else if ch == '\\' {
			isEscaped = true
			isSeparated = false
			continue
		}
		if isSeparated {
			args = append(args, "")
		}
		current := &args[len(args)-1]
		if !isEscaped && ch == '"' {
			i++
			for ; i < len(src) && !(!isEscaped && src[i] == '"'); i++ {
				ch = src[i]
				if ch == '\\' {
					isEscaped = true
				} else {
					*current += string(ch)
					isEscaped = false
				}
			}
			isEscaped = false
			isSeparated = false
			continue
		}
		*current += string(ch)
		isEscaped = false
		isSeparated = false
	}
	return args, nil
}

// UnescapeString convert '\\' to single '\', '\n' to new line, '\t', '\"' to double quote to tab to new line
func UnescapeString(src string) string {
	var (
		buf     []byte
		c       byte
		in, out int
	)
	if !strings.Contains(src, "\\") {
		return src
	}
	buf = make([]byte, len(src))
	for in = 0; in < len(src); in++ {
		c = src[in]
		if c != '\\' {
			buf[out] = c
			out++
			continue
		}
		in++
		c = src[in]
		switch c {
		case 'n':
			buf[out] = '\n'
		case 't':
			buf[out] = '\t'
		case '"':
			buf[out] = '"'
		case '\'':
			buf[out] = '\''
		default:
			buf[out] = c
		}
		out++
	}
	return string(buf[:out])
}
