package dependency

import (
	"fmt"
	"reflect"
)

// SubProvider is default sub-dependency provider
type SubProvider struct {
	parent Provider
	pool   map[string]*Builder
}

// NewSubProvider create new instance of a sub-depenedency provider
func NewSubProvider(p Provider) Provider {
	return Provider(&SubProvider{
		parent: p,
		pool:   map[string]*Builder{},
	})
}

// Get return instance by name
func (d *SubProvider) Get(faceName string) (Instance, error) {
	if row, exist := d.pool[faceName]; exist {
		instance, err := row.Get(d)
		if err != nil {
			return nil, err
		}
		return instance, nil
	}
	return d.parent.Get(faceName)
}

// InjectTo inject dependencies to object
func (d *SubProvider) InjectTo(obj interface{}) error {
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

// GetAll return a map of all dependency builders
func (d *SubProvider) GetAll() map[string]*Builder {
	tmpMap := make(map[string]*Builder)
	for key, val := range d.parent.GetAll() {
		tmpMap[key] = val
	}
	for key, val := range d.pool {
		tmpMap[key] = val
	}
	return tmpMap
}

// AddService add new named service
func (d *SubProvider) AddService(name string, factory Factory) error {
	return d.add(name, factory, true, false)
}

// AddDefaultService add default service (it can be overwrite by AddService)
func (d *SubProvider) AddDefaultService(name string, factory Factory) error {
	return d.addDefault(name, factory, true)
}

func (d *SubProvider) addDefault(name string, factory Factory, static bool) error {
	if _, exist := d.pool[name]; exist {
		return fmt.Errorf("default can be set once")
	}
	return d.add(name, factory, static, true)
}

func (d *SubProvider) add(name string, factory Factory, static, isDefault bool) error {
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
