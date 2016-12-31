package provider

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/goatcms/goat-core/dependency"
)

// Provider is default dependency distributor
type Provider struct {
	defaults  map[string]dependency.Factory
	factories map[string]dependency.Factory
	instances map[string]interface{}
	callstack []string
	keys      []string
	blocked   bool
	tagname   string
}

// NewProvider create new instance of a depenedency provider
func NewProvider(tagname string) dependency.Provider {
	return &Provider{
		defaults:  map[string]dependency.Factory{},
		factories: map[string]dependency.Factory{},
		instances: map[string]interface{}{},
		callstack: []string{},
		keys:      []string{},
		blocked:   false,
		tagname:   tagname,
	}
}

// NewStaticProvider create a dependency provider from Factories map. It is static (mean that it is pre-defined and blocked for modifications)
func NewStaticProvider(tagname string, factories map[string]dependency.Factory) dependency.Provider {
	keys := make([]string, len(factories))
	i := 0
	for key, _ := range factories {
		keys[i] = key
		i++
	}
	return &Provider{
		defaults:  map[string]dependency.Factory{},
		factories: map[string]dependency.Factory{},
		instances: map[string]interface{}{},
		callstack: []string{},
		keys:      keys,
		blocked:   true,
		tagname:   tagname,
	}
}

// Keys return list of all defined dependencies names
func (d *Provider) Keys() ([]string, error) {
	return d.keys, nil
}

// Block prevent nev dependency definition
func (d *Provider) Block() {
	d.blocked = true
}

// Get return instance by name
func (d *Provider) Get(name string) (interface{}, error) {
	d.blocked = true
	if d.isCalled(name) {
		return nil, fmt.Errorf("%s is cyclic dependency (dependency callstack: %v)", name, append(d.callstack, name))
	}
	if instance, exist := d.instances[name]; exist {
		return instance, nil
	}
	if factory, exist := d.factories[name]; exist {
		d.callstack = append(d.callstack, name)
		instance, err := factory(d)
		if err != nil {
			return nil, fmt.Errorf("%v (dependency callstack: %v)", err, d.callstack)
		}
		if instance == nil {
			return nil, fmt.Errorf("factory for %s return nil as instance", name)
		}
		d.callstack = d.callstack[:len(d.callstack)-1]
		d.clean(name)
		d.instances[name] = instance
		return instance, nil
	}
	if factory, exist := d.defaults[name]; exist {
		d.callstack = append(d.callstack, name)
		instance, err := factory(d)
		if err != nil {
			return nil, fmt.Errorf("%v (dependency callstack: %v)", err, d.callstack)
		}
		if instance == nil {
			return nil, fmt.Errorf("default factory for %s return nil as instance", name)
		}
		d.callstack = d.callstack[:len(d.callstack)-1]
		delete(d.defaults, name)
		d.instances[name] = instance
		return instance, nil
	}
	return nil, fmt.Errorf("goatcore/dependency/provider: dependency %s doesn't exist", name)
}

// Set define new instance
func (d *Provider) Set(name string, instance interface{}) error {
	if d.blocked {
		return fmt.Errorf("goatcore/dependency/provider: can not add new instance after first get dependency (for %s)", name)
	}
	if _, exist := d.instances[name]; exist {
		return fmt.Errorf("goatcore/dependency/provider: dependency %s exist (and musn't be overwrited)", name)
	}
	if _, exist := d.factories[name]; exist {
		return fmt.Errorf("goatcore/dependency/provider: dependency %s factory exist (and value musn't be overwrited)", name)
	}
	d.keys = append(d.keys, name)
	d.instances[name] = instance
	d.addKey(name)
	return nil
}

// AddFactory define a factory for dependency
func (d *Provider) AddFactory(name string, factory dependency.Factory) error {
	if d.blocked {
		return fmt.Errorf("goatcore/dependency/provider: can not add factory after first get dependency (for %s)", name)
	}
	if _, exist := d.factories[name]; exist {
		return fmt.Errorf("goatcore/dependency/provider: factory for '%s' double defined", name)
	}
	d.clean(name)
	d.factories[name] = factory
	d.addKey(name)
	return nil
}

// AddDefaultFactory define a default factory for dependency
func (d *Provider) AddDefaultFactory(name string, factory dependency.Factory) error {
	if d.blocked {
		return fmt.Errorf("goatcore/dependency/provider: can not add factory after first get dependency (for %s)", name)
	}
	if _, exist := d.defaults[name]; exist {
		return fmt.Errorf("goatcore/dependency/provider: default factory for '%s' double defined", name)
	}
	if _, exist := d.factories[name]; exist {
		// when we have got defined factory for a field default factory wont be used
		return nil
	}
	d.defaults[name] = factory
	d.addKey(name)
	return nil
}

// InjectTo inject dependencies to object
func (d *Provider) InjectTo(obj interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	for i := 0; i < structValue.NumField(); i++ {
		var isRequired = true

		valueField := structValue.Field(i)
		structField := structValue.Type().Field(i)

		depID := structField.Tag.Get(d.tagname)
		if depID == "" {
			continue
		}
		if strings.HasPrefix(depID, "?") {
			isRequired = false
			depID = depID[1:]
		}
		if !valueField.IsValid() {
			return fmt.Errorf("goatcore/dependency/provider.InjectTo: %s is invalid", structField.Name)
		}
		if !valueField.CanSet() {
			return fmt.Errorf("goatcore/dependency/provider.InjectTo: Cannot set %s field value", structField.Name)
		}
		dep, err := d.Get(depID)
		if err != nil {
			if !isRequired {
				continue
			}
			return err
		}
		if dep == nil {
			return fmt.Errorf("goatcore/dependency/provider.InjectTo: dependency instance can not be nil (%s)", depID)
		}
		depValue := reflect.ValueOf(dep)
		valueField.Set(depValue)
	}
	return nil
}

func (d *Provider) clean(name string) {
	if _, exist := d.factories[name]; exist {
		delete(d.factories, name)
	}
	if _, exist := d.defaults[name]; exist {
		delete(d.defaults, name)
	}
}

func (d *Provider) addKey(name string) {
	if !d.hasKey(name) {
		d.keys = append(d.keys, name)
	}
}

func (d *Provider) isCalled(name string) bool {
	for _, v := range d.callstack {
		if v == name {
			return true
		}
	}
	return false
}

func (d *Provider) hasKey(name string) bool {
	for _, v := range d.keys {
		if v == name {
			return true
		}
	}
	return false
}
