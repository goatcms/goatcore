package namespaces

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/dependency"
)

const (
	namespacePrefix = "namespaces"
)

// Unit is namespace storage for scope
type Unit struct{}

// NewUnit create new Unit instance
func NewUnit() Unit {
	return Unit{}
}

// UnitFactory create new Unit instance
func UnitFactory(dp dependency.Provider) (ri interface{}, err error) {
	return pipservices.NamespacesUnit(NewUnit()), nil
}

// FromScope return scope namespaces
func (unit Unit) FromScope(scp app.Scope, defaultNamespace pipservices.Namespaces) (namespace pipservices.Namespaces, err error) {
	var ins interface{}
	if ins, err = scp.Get(scopeKey); err != nil {
		return nil, err
	}
	if ins == nil {
		return defaultNamespace, nil
	}
	return ins.(pipservices.Namespaces), nil
}

// Define define scope namespaces
func (unit Unit) Define(scp app.Scope, namespaces pipservices.Namespaces) (err error) {
	return scp.Set(scopeKey, namespaces)
}

// Bind set child scope namespace from parent
func (unit Unit) Bind(parent, child app.Scope) (err error) {
	var namespace pipservices.Namespaces
	if namespace, err = unit.FromScope(parent, DefaultNamespace); err != nil {
		return err
	}
	return child.Set(scopeKey, namespace)
}
