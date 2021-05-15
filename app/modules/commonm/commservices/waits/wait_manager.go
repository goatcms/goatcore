package waits

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
)

// WaitManager is namespace storage for scope
type WaitManager struct{}

// NewWaitManager create new WaitManager instance
func NewWaitManager() WaitManager {
	return WaitManager{}
}

// WaitManagerFactory create new WaitManager instance
func WaitManagerFactory(dp app.DependencyProvider) (ri interface{}, err error) {
	return commservices.WaitManager(NewWaitManager()), nil
}

// ForScope return scope manager
func (manager WaitManager) ForScope(ctxScope app.Scope) (swm commservices.ScopeWaitManager, err error) {
	var (
		locker app.DataScopeLocker
		v      interface{}
	)
	locker = ctxScope.LockData()
	v = locker.Value(waitManagerKey)
	if v == nil {
		swm = NewScopeWaitManager(ctxScope)
		locker.SetValue(waitManagerKey, swm)
	} else {
		swm = v.(commservices.ScopeWaitManager)
	}
	locker.Commit()
	return swm, nil
}
