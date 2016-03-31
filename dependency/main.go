package dependency

type Instance interface{}
type Factory func(Provider) (Instance, error)

type Provider interface {
	Get(string) (Instance, error)
	AddService(string, Factory) error
	AddFactory(string, Factory) error
	AddDefaultService(string, Factory) error
	AddDefaultFactory(string, Factory) error
}

type Loadable interface {
	Load(*Provider) error
}
