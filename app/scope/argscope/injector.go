package argscope

import (
	"strconv"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil"
)

// InjectString insert arguments from string to data scope
func InjectString(scp app.DataScope, str string) (err error) {
	var args []string
	if args, _, err = varutil.SplitArguments(str); err != nil {
		return
	}
	return InjectArgs(scp, args...)
}

// InjectArgs insert arguments to data scope
func InjectArgs(scp app.DataScope, args ...string) (err error) {
	// remove params after "--" without name
	// it is popular command separator
	var separated []string
	args, separated = SeparateArgs(args)
	scp.SetValue("--", separated)
	// scan arguments
	anonIndex := 0
	for _, arg := range args {
		if strings.Contains(arg, "=") {
			// named argument
			arg = strings.TrimPrefix(arg, "-") // reduce first "-" (if required)
			arg = strings.TrimPrefix(arg, "-") // reduce second "-" (if required)
			key, value := separate(arg, "true")
			scp.SetValue(key, value)
			continue
		}
		// anonymous argument
		key := "$" + strconv.Itoa(anonIndex)
		scp.SetValue(key, arg)
		anonIndex++
	}
	return nil
}

func separate(arg, defaultValue string) (key, value string) {
	index := strings.Index(arg, "=")
	if index == -1 {
		return arg, defaultValue
	}
	return arg[:index], arg[index+1:]
}
