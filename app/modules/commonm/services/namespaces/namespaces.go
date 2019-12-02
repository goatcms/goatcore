package namespaces

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/services"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

const (
	namespacePrefix = "namespaces"
)

// Namasepaces is namespace storage for scope
type Namasepaces struct{}

// NewNamasepaces create new Namasepaces instance
func NewNamasepaces() Namasepaces {
	return Namasepaces{}
}

// NamasepacesFactory create new Namasepaces instance
func NamasepacesFactory(dp dependency.Provider) (ri interface{}, err error) {
	return services.Namespaces(NewNamasepaces()), nil
}

// Set define new scope namespace
func (namasepaces Namasepaces) Set(ctxScope app.Scope, name, value string) (err error) {
	var (
		key = namespacePrefix + name
		v   interface{}
	)
	if v, err = ctxScope.Get(key); err == nil || v != nil {
		return goaterr.Errorf("Namasepaces.Set: Namespace can not be defined twice")
	}
	return ctxScope.Set(key, value)
}

// Get return pip by name
func (namasepaces Namasepaces) Get(ctxScope app.Scope, name string) (value string, err error) {
	var key = namespacePrefix + name
	return scope.GetString(ctxScope, key)
}
