package definition

import (
	"strings"
)

type Params []string

func (p Params) Args() []string {
	return p
}

func (p Params) Key(name string) string {
	prefix := name + ":"
	for _, v := range p {
		if strings.HasPrefix(v, prefix) {
			return v[len(prefix):]
		}
	}
	return ""
}
