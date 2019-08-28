package scopesync

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil/goaterr"
	"github.com/goatcms/goatcore/workers/jobsync"
)

// AppendError add error to scope
func AppendError(scope app.Scope, err error) {
	var lifecycle *jobsync.Lifecycle
	if err == nil {
		return
	}
	lifecycle = Lifecycle(scope)
	lifecycle.Error(err)
	scope.Trigger(app.ErrorEvent, err)
}

// ToError return scope error object or nil if scope has no error
func ToError(scope app.Scope) error {
	var lifecycle = Lifecycle(scope)
	return goaterr.ToErrors(lifecycle.Errors())
}
