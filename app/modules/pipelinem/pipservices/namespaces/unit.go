package namespaces

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
)

// Unit is namespace storage for scope
type Unit struct{}

// NewUnit create new Unit instance
func NewUnit() Unit {
	return Unit{}
}

// UnitFactory create new Unit instance
func UnitFactory(dp app.DependencyProvider) (ri interface{}, err error) {
	return pipservices.NamespacesUnit(NewUnit()), nil
}

// FromScope return scope namespaces
func (unit Unit) FromScope(scp app.Scope, defaultNamespace pipservices.Namespaces) (namespace pipservices.Namespaces, err error) {
	ins := scp.Value(scopeKey)
	if ins == nil {
		return defaultNamespace, nil
	}
	return ins.(pipservices.Namespaces), nil
}

// Define define scope namespaces
func (unit Unit) Define(scp app.Scope, namespaces pipservices.Namespaces) (err error) {
	scp.SetValue(scopeKey, namespaces)
	return nil
}

// Bind set child scope namespace from parent
func (unit Unit) Bind(parent, child app.Scope) (err error) {
	var namespace pipservices.Namespaces
	if namespace, err = unit.FromScope(parent, DefaultNamespace); err != nil {
		return err
	}
	child.SetValue(scopeKey, namespace)
	return nil
}
