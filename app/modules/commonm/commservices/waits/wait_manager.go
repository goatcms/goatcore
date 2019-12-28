package waits

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/dependency"
)

// WaitManager is namespace storage for scope
type WaitManager struct{}

// NewWaitManager create new WaitManager instance
func NewWaitManager() WaitManager {
	return WaitManager{}
}

// WaitManagerFactory create new WaitManager instance
func WaitManagerFactory(dp dependency.Provider) (ri interface{}, err error) {
	return commservices.WaitManager(NewWaitManager()), nil
}

// ForScope return scope manager
func (manager WaitManager) ForScope(ctxScope app.Scope) (swm commservices.ScopeWaitManager, err error) {
	var (
		locker app.DataScopeLocker
		v      interface{}
	)
	locker = ctxScope.LockData()
	defer locker.Commit()
	if v, err = ctxScope.Get(waitManagerKey); err != nil {
		return nil, err
	}
	if v == nil {
		swm = NewScopeWaitManager(ctxScope)
		if err = ctxScope.Set(waitManagerKey, swm); err != nil {
			return nil, err
		}
	} else {
		swm = v.(commservices.ScopeWaitManager)
	}
	return swm, nil
}
