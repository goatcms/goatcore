package goatapp

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/dependency"
)

// DependencyScope represent dependency scope lvl
type DependencyScope struct {
	app.EventScope
	dependency.Provider
}

// NewDependencyScope create new dependency scope
func NewDependencyScope(dp dependency.Provider) app.Scope {
	return DependencyScope{
		EventScope: scope.NewEventScope(),
		Provider:   dp,
	}
}
