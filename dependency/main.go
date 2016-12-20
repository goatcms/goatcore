package dependency

// Instance represent instance of a service
type Instance interface{}

// Factory represent a builder of a dependency instance
type Factory func(Provider) (Instance, error)

// Provider distribute dependencies
type Provider interface {
	Get(string) (Instance, error)
	Set(string, Instance) error
	Keys() []string
	AddFactory(string, Factory) error
	AddDefaultFactory(string, Factory) error
	InjectTo(obj interface{}) error
}
