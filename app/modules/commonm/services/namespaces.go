package services

import "github.com/goatcms/goatcore/app"

// Namespaces storage context namespaces
type Namespaces interface {
	// Set new scope namespace by name
	Set(scope app.Scope, namespace, value string) (err error)
	// Get return scope namespace by name
	Get(scope app.Scope, namespace string) (value string, err error)
}
