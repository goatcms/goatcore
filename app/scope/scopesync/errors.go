package scopesync

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// AppendError add error to scope
func AppendError(scope app.Scope, err error) {
	var lifecycle = Lifecycle(scope)
	lifecycle.Error(err)
	lifecycle.Kill()
	scope.Trigger(app.ErrorEvent, err)
}

// ToError return scope error object or nil if scope has no error
func ToError(scope app.Scope) error {
	var lifecycle = Lifecycle(scope)
	return goaterr.ToErrors(lifecycle.Errors())
}
