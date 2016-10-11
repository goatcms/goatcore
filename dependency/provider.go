package dependency

import (
	"fmt"
	"reflect"
)

// DefaultProvider is default dependency distributor
type DefaultProvider struct {
	pool map[string]*Builder
}

// NewProvider create new instance of a depenedency provider
func NewProvider() Provider {
	return &DefaultProvider{
		pool: map[string]*Builder{},
	}
}

// Get return instance by name
func (d *DefaultProvider) Get(faceName string) (Instance, error) {
	if row, exist := d.pool[faceName]; exist {
		instance, err := row.Get(d)
		if err != nil {
			return nil, err
		}
		return instance, nil
	}
	return nil, fmt.Errorf("implementation of interface %v not exist", faceName)
}

// GetAll return a map of all dependency builders
func (d *DefaultProvider) GetAll() map[string]*Builder {
	return d.pool
}

// AddService add new named service
func (d *DefaultProvider) AddService(name string, factory Factory) error {
	return d.add(name, factory, true, false)
}

/*func (d *DefaultProvider) AddFactory(name string, factory Factory) error {
	return d.add(name, factory, false, false)
}*/

// AddDefaultService add default service (it can be overwrite by AddService)
func (d *DefaultProvider) AddDefaultService(name string, factory Factory) error {
	return d.addDefault(name, factory, true)
}

// InjectTo inject dependencies to object
func (d *DefaultProvider) InjectTo(obj interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	for i := 0; i < structValue.NumField(); i++ {
		valueField := structValue.Field(i)
		structField := structValue.Type().Field(i)

		depID := structField.Tag.Get(InjectTagName)
		if depID == "" {
			continue
		}
		if !valueField.IsValid() {
			return fmt.Errorf("DependencyProvider.InjectTo: %s is not valid", structField.Name)
		}
		if !valueField.CanSet() {
			return fmt.Errorf("DependencyProvider.InjectTo: Cannot set %s field value", structField.Name)
		}
		dep, err := d.Get(depID)
		if err != nil {
			return err
		}
		if dep == nil {
			return fmt.Errorf("DependencyProvider.InjectTo: dependency instance can not be nil (%s)", depID)
		}
		depValue := reflect.ValueOf(dep)
		valueField.Set(depValue)
	}
	return nil
}

/*func (d *DefaultProvider) AddDefaultFactory(name string, factory Factory) error {
	return d.addDefault(name, factory, false)
}*/

func (d *DefaultProvider) addDefault(name string, factory Factory, static bool) error {
	if _, exist := d.pool[name]; exist {
		return fmt.Errorf("default can be set once")
	}
	return d.add(name, factory, static, true)
}

func (d *DefaultProvider) add(name string, factory Factory, static, isDefault bool) error {
	if current, exist := d.pool[name]; exist {
		if !current.isDefault {
			return fmt.Errorf("interface %v exist", name)
		}
	}
	d.pool[name] = &Builder{
		Static:    static,
		Factory:   factory,
		Instance:  nil,
		isCalled:  false,
		isDefault: isDefault,
	}
	return nil
}
