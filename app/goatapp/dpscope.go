package goatapp

import (
	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goat-core/app/scope"
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goat-core/dependency/provider"
)

// DependencyScope represent dependency scope lvl
type DependencyScope struct {
	app.EventScope
	DP dependency.Provider
}

// NewDependencyScope create new dependency scope
func NewDependencyScope(tagname string) app.Scope {
	return DependencyScope{
		EventScope: scope.NewEventScope(),
		DP:         provider.NewProvider(tagname),
	}
}

// Set is deprecated. It exists to interface support.
func (ds DependencyScope) Set(name string, instance interface{}) error {
	return ds.DP.Set(name, instance)
}

// Get dependency by name
func (ds DependencyScope) Get(key string) interface{} {
	instance, err := ds.DP.Get(key)
	if err != nil {
		return nil
	}
	return instance
}

// Keys return list of keys
func (ds DependencyScope) Keys() []string {
	return ds.DP.Keys()
}

// InjectTo inject data to object
func (ds DependencyScope) InjectTo(obj interface{}) error {
	return ds.DP.InjectTo(obj)
}
