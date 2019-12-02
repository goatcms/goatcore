package argscope

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/varutil"
)

// NewScope create new ArgScope instance
func NewScope(args []string, tagname string) (ins app.Scope, err error) {
	ins = scope.NewScope(tagname)
	if err = InjectArgsToScope(args, ins); err != nil {
		return nil, err
	}
	return ins, nil
}

// NewScopeFromString create argument scope from string
func NewScopeFromString(line, tagname string) (s app.Scope, err error) {
	var args []string
	if args, _, err = varutil.SplitArguments(line); err != nil {
		return nil, err
	}
	return NewScope(args, tagname)
}
