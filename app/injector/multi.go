package injector

import "github.com/goatcms/goat-core/app"

// MultiInjector is wrap for many injection interface to run them by onc call
type MultiInjector struct {
	injectors []app.Injector
}

// NewMultiInjector create new multi injection instance
func NewMultiInjector(injectors []app.Injector) app.Injector {
	return app.Injector(MultiInjector{
		injectors: injectors,
	})
}

// InjectTo inject data from all injectors
func (mi MultiInjector) InjectTo(obj interface{}) error {
	for _, injector := range mi.injectors {
		if err := injector.InjectTo(obj); err != nil {
			return err
		}
	}
	return nil
}
