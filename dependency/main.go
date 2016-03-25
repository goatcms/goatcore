package dependency

type Instance interface{}
type Factory func() (Instance, error)

type Provider interface {
	Get(string) (Instance, error)
	AddService(string, Factory) error
	AddFactory(string, Factory) error
}

type Loadable interface {
	Load(*Provider) error
}
