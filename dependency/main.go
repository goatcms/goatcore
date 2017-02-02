package dependency

// Factory represent a builder of a dependency instance
type Factory func(Provider) (interface{}, error)

// Provider distribute dependencies
type Provider interface {
	Injector
	AddInjectors([]Injector) error
	Get(string) (interface{}, error)
	Set(string, interface{}) error
	Keys() ([]string, error)
	SetDefault(string, interface{}) error
	AddFactory(string, Factory) error
	AddDefaultFactory(string, Factory) error
}

// Injector provide interface to inject data
type Injector interface {
	InjectTo(obj interface{}) error
}
