package varutil

import (
	"io"
	"strings"

	"github.com/goatcms/goatcore/varutil/goaterr"
)

// SplitArguments convert string to separated arguments array
func SplitArguments(src string) (args []string, eof bool, err error) {
	return ReadArguments(strings.NewReader(src))
}

// ReadArguments convert string to separated arguments array
func ReadArguments(reader io.Reader) (args []string, eof bool, err error) {
	var (
		isEscaped   = false
		isSeparated = true
		buf         = make([]byte, 1)
		ch          rune
	)
	for {
		if _, err = reader.Read(buf); err != nil {
			if err == io.EOF {
				return args, true, nil
			}
			return nil, false, err
		}
		ch = rune(buf[0])
		if ch == '\n' {
			if isEscaped {
				isEscaped = false
				continue
			}
			return args, false, nil
		}
		if ch == ' ' || ch == '\t' {
			isEscaped = false
			isSeparated = true
			continue
		} else if !isEscaped && ch == '\\' {
			isEscaped = true
			isSeparated = false
			continue
		}
		if isSeparated {
			args = append(args, "")
		}
		current := &args[len(args)-1]
		if !isEscaped && ch == '"' {
			for {
				if _, err = reader.Read(buf); err != nil {
					return nil, err == io.EOF, err
				}
				ch = rune(buf[0])
				if !isEscaped && ch == '"' {
					break
				}
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
		if !isEscaped && ch == '<' && strings.HasSuffix(*current, "=<") {
			var (
				eof   = ""
				base  = *current
				value = ""
			)
			base = base[:len(base)-1]
			// read escape sequence
			for {
				if _, err = reader.Read(buf); err != nil {
					return nil, err == io.EOF, goaterr.Errorf(err.Error())
				}
				ch = rune(buf[0])
				if ch == '\n' {
					break
				}
				if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_' {
					eof += string(ch)
				} else if ch != ' ' && ch != '\t' {
					return nil, false, goaterr.Errorf("argument EOF sequence of multiline value can include only low and upper letters ")
				}
			}
			if eof == "" {
				return nil, false, goaterr.Errorf("insert EOF sequence after open multiline argument by '=<<' sequence'")
			}
			eof = "\n" + eof
			// read data
			for {
				if _, err = reader.Read(buf); err != nil {
					return nil, err == io.EOF, goaterr.Errorf(err.Error())
				}
				value += string(buf[0])
				if strings.HasSuffix(value, eof) {
					value = value[:len(value)-len(eof)]
					break
				}
			}
			value = base + strings.Trim(value, " \t")
			args[len(args)-1] = value
			continue
		}
		*current += string(ch)
		isEscaped = false
		isSeparated = false
	}
}
