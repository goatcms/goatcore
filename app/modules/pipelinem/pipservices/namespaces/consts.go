package namespaces

import "github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"

const (
	scopeKey = "pipNamespaces"
)

var (
	// DefaultNamespace is default namespace for main task
	DefaultNamespace = NewNamespaces(pipservices.NamasepacesParams{
		Task: "main",
		Lock: "",
	})
)
