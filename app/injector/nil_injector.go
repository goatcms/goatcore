package injector

import (
	"github.com/goatcms/goatcore/app"
)

var (
	nilInjectorInstance = NilInjector{}
)

// NilInjector is empty injector
type NilInjector struct{}

// NewNilInjector create new map injector instance
func NewNilInjector() app.Injector {
	return nilInjectorInstance
}

// InjectTo do nothing
func (injector NilInjector) InjectTo(obj interface{}) error {
	return nil
}
