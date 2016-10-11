package dependency

const (
	// InjectTagName is name use for extract dependency name from tag
	InjectTagName = "inject"
	// SelfName is name to inject itsself
	SelfName = "provider"
)

// Instance represent instance of a service
type Instance interface{}

// Factory represent a builder of a dependency instance
type Factory func(Provider) (Instance, error)

// Provider distribute dependencies
type Provider interface {
	Get(string) (Instance, error)
	GetAll() map[string]*Builder
	AddService(string, Factory) error
	//AddFactory(string, Factory) error
	AddDefaultService(string, Factory) error
	//AddDefaultFactory(string, Factory) error
	InjectTo(obj interface{}) error
}

// Loadable represents loadable interface
type Loadable interface {
	Load(*Provider) error
}
