package scopesync

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/workers"
	"github.com/goatcms/goatcore/workers/jobsync"
)

const (
	insKey = "_scopesync.lifecycle"
)

func Lifecycle(scope app.Scope) *jobsync.Lifecycle {
	var (
		lifecycleIns interface{}
		lifecycle    *jobsync.Lifecycle
		err          error
	)
	lifecycleIns, err = scope.Get(insKey)
	if err != nil || lifecycleIns == nil {
		lifecycle = jobsync.NewLifecycle(workers.DefaultTimeout, true)
		scope.Set(insKey, lifecycle)
	} else {
		lifecycle = lifecycleIns.(*jobsync.Lifecycle)
	}
	return lifecycle
}
