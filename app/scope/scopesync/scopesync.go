package scopesync

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/workers/jobsync"
)

// Lifecycle return jobsync.Lifecycle object connected to scope. If Lifecycle object doesn't exist, create new one and return it.
func Lifecycle(scope app.Scope) *jobsync.Lifecycle {
	var (
		lifecycleIns interface{}
		lifecycle    *jobsync.Lifecycle
		err          error
	)
	lifecycleIns, err = scope.Get(insKey)
	if err != nil || lifecycleIns == nil {
		lifecycle = jobsync.NewLifecycle(app.DefaultDeadline, true)
		scope.Set(insKey, lifecycle)
	} else {
		lifecycle = lifecycleIns.(*jobsync.Lifecycle)
	}
	return lifecycle
}
