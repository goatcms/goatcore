package app

// Injector inject data/dependencies to object
type Injector interface {
	InjectTo(obj interface{}) error
}
