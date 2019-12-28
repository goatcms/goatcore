package namespaces

const (
	scopeKey = "pipNamespaces"
)

var (
	// DefaultNamespace is default namespace for main task
	DefaultNamespace = NewNamespaces("main", "")
)
