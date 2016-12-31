package dependency

// Factory represent a builder of a dependency instance
type Factory func(Provider) (interface{}, error)

// Provider distribute dependencies
type Provider interface {
	Get(string) (interface{}, error)
	Set(string, interface{}) error
	Keys() ([]string, error)
	AddFactory(string, Factory) error
	AddDefaultFactory(string, Factory) error
	InjectTo(obj interface{}) error
}
