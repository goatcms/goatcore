package pipservices

import "github.com/goatcms/goatcore/app"

// Namespaces storage namespaces
type Namespaces interface {
	// Task return task namespace
	Task() (value string)
	// Lock return lock namespace
	Lock() (value string)
}

// NamespacesUnit is a service to controll namespaces
type NamespacesUnit interface {
	// FromScope return namespace from scope
	FromScope(scp app.Scope, defaultNamespace Namespaces) (namespace Namespaces, err error)
	// DefineScope define scope namespaces
	Define(scp app.Scope, task, lock string) (err error)
	// Bind copy namespaces from parent scope to child scope
	Bind(parent, child app.Scope) (err error)
}
