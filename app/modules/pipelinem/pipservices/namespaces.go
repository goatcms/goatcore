package pipservices

import "github.com/goatcms/goatcore/app"

// NamasepacesParams is params for scope
type NamasepacesParams struct {
	Task string
	Lock string
}

// Namespaces storage namespaces
type Namespaces interface {
	// Task return task namespace
	Task() (value string)
	// Lock return lock namespace
	Lock() (value string)
}

// NamespacesUnit is a service to controll namespaces
type NamespacesUnit interface {
	// FromScope return namespace from scope (or defaultNamespace if scope is undefined)
	FromScope(scp app.Scope, defaultNamespace Namespaces) (namespace Namespaces, err error)
	// DefineScope define scope namespaces
	Define(scp app.Scope, namspaces Namespaces) (err error)
	// Bind copy namespaces from parent scope to child scope
	Bind(parent, child app.Scope) (err error)
}
