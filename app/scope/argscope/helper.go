package argscope

import (
	"strconv"
	"strings"

	"github.com/goatcms/goatcore/app"
)

// InjectArgsToScope insert arguments to scope
func InjectArgsToScope(args []string, scope app.DataScope) (err error) {
	for i, value := range args {
		// position keys
		ikey := "$" + strconv.Itoa(i)
		scope.SetValue(ikey, value)
		// reduce prefixes
		if strings.HasPrefix(value, "--") {
			value = value[2:]
		} else if strings.HasPrefix(value, "-") {
			value = value[1:]
		}
		// key:value
		index := strings.Index(value, "=")
		if index != -1 {
			key := value[:index]
			value = value[index+1:]
			scope.SetValue(key, value)
		} else {
			scope.SetValue(value, "true")
		}
	}
	return nil
}
