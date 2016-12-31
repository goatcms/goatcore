package goatapp

import (
	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goat-core/app/scope"
	"github.com/goatcms/goat-core/dependency"
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
