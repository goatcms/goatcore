package commservices

import (
	"github.com/goatcms/goatcore/app"
)

// WaitManager provide waitgroups
type WaitManager interface {
	// ForScope return ScopeWaitManager for scope
	ForScope(s app.Scope) (swm ScopeWaitManager, err error)
}

// ScopeWaitManager provide waitgroups
type ScopeWaitManager interface {
	Add(name string, count int)
	Done(name string)
	Wait(name string)
}
