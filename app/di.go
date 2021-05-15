package app

// Factory represent a builder of a dependency instance
type Factory func(DependencyProvider) (interface{}, error)

// Provider distribute dependencies
type DependencyProvider interface {
	Injector
	AddInjectors([]Injector) error
	Get(string) (interface{}, error)
	Set(string, interface{}) error
	Keys() ([]string, error)
	SetDefault(string, interface{}) error
	AddFactory(string, Factory) error
	AddDefaultFactory(string, Factory) error
}

// Injector inject data/dependencies to object
type Injector interface {
	InjectTo(obj interface{}) error
}
