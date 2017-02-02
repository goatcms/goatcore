package provider

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/goatcms/goat-core/dependency"
)

// Provider is default dependency distributor
type Provider struct {
	injectors        []dependency.Injector
	defaultFactories map[string]dependency.Factory
	factories        map[string]dependency.Factory
	defaultInstances map[string]interface{}
	instances        map[string]interface{}
	callstack        []string
	keys             []string
	blocked          bool
	autoclean        bool
	tagname          string
}

// NewProvider create new instance of a depenedency provider
func NewProvider(tagname string) dependency.Provider {
	return &Provider{
		injectors:        []dependency.Injector{},
		defaultFactories: map[string]dependency.Factory{},
		factories:        map[string]dependency.Factory{},
		defaultInstances: map[string]interface{}{},
		instances:        map[string]interface{}{},
		callstack:        []string{},
		keys:             []string{},
		blocked:          false,
		autoclean:        true,
		tagname:          tagname,
	}
}

// NewStaticProvider create a dependency provider from Factories map. It is static (mean that it is pre-defined and blocked for modifications)
func NewStaticProvider(tagname string, factories map[string]dependency.Factory, instances map[string]interface{}, injectors []dependency.Injector) dependency.Provider {
	keys := make([]string, len(factories))
	i := 0
	for key, _ := range factories {
		keys[i] = key
		i++
	}
	return &Provider{
		injectors:        injectors,
		defaultFactories: map[string]dependency.Factory{},
		factories:        factories,
		defaultInstances: map[string]interface{}{},
		instances:        instances,
		callstack:        []string{},
		keys:             keys,
		blocked:          true,
		autoclean:        false,
		tagname:          tagname,
	}
}

// AddInjectors add new injector to dependency provider
func (d *Provider) AddInjectors(injectors []dependency.Injector) error {
	if d.blocked {
		return fmt.Errorf("goatcore/dependency/provider.AddInjectors: can not add new injector after got dependency")
	}
	d.injectors = append(d.injectors, injectors...)
	return nil
}

// Keys return list of all defined dependencies names
func (d *Provider) Keys() ([]string, error) {
	return d.keys, nil
}

// Block prevent nev dependency definition
func (d *Provider) Block() {
	if d.blocked == true {
		return
	}
	for key, defaultVal := range d.defaultInstances {
		if _, ok := d.instances[key]; !ok {
			if d.autoclean {
				delete(d.defaultFactories, key)
				delete(d.factories, key)
			}
			d.instances[key] = defaultVal
		}
	}
	d.defaultInstances = nil
	d.blocked = true
}

// Get return instance by name
func (d *Provider) Get(name string) (interface{}, error) {
	d.Block()
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
	if factory, exist := d.defaultFactories[name]; exist {
		d.callstack = append(d.callstack, name)
		instance, err := factory(d)
		if err != nil {
			return nil, fmt.Errorf("%v (dependency callstack: %v)", err, d.callstack)
		}
		if instance == nil {
			return nil, fmt.Errorf("default factory for %s return nil as instance", name)
		}
		d.callstack = d.callstack[:len(d.callstack)-1]
		if d.autoclean {
			delete(d.defaultFactories, name)
		}
		d.instances[name] = instance
		return instance, nil
	}
	return nil, fmt.Errorf("goatcore/dependency/provider: dependency %s doesn't exist", name)
}

// Set instance
func (d *Provider) Set(name string, instance interface{}) error {
	if d.blocked {
		return fmt.Errorf("goatcore/dependency/provider.Set: can not add new instance after got dependency (for %s)", name)
	}
	if _, exist := d.instances[name]; exist {
		return fmt.Errorf("goatcore/dependency/provider.Set: dependency %s exist (musn't be overwrited)", name)
	}
	if _, exist := d.factories[name]; exist {
		return fmt.Errorf("goatcore/dependency/provider.Set: dependency %s factory exists (value musn't be overwrited)", name)
	}
	d.instances[name] = instance
	d.addKey(name)
	return nil
}

// Set defult instance
func (d *Provider) SetDefault(name string, instance interface{}) error {
	if d.blocked {
		return fmt.Errorf("goatcore/dependency/provider.SetDefault: can not add default instance after got dependency (for %s)", name)
	}
	if _, exist := d.defaultInstances[name]; exist {
		return fmt.Errorf("goatcore/dependency/provider.SetDefault: dependency %s exist (musn't be overwrited)", name)
	}
	if _, exist := d.defaultFactories[name]; exist {
		return fmt.Errorf("goatcore/dependency/provider.SetDefault: dependency %s factory exists (musn't be overwrited)", name)
	}
	d.defaultInstances[name] = instance
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
	if _, exist := d.defaultFactories[name]; exist {
		return fmt.Errorf("goatcore/dependency/provider: default factory for '%s' double defined", name)
	}
	if _, exist := d.factories[name]; exist {
		// when we have got defined factory for a field default factory wont be used
		return nil
	}
	d.defaultFactories[name] = factory
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
	for _, injector := range d.injectors {
		if err := injector.InjectTo(obj); err != nil {
			return err
		}
	}
	return nil
}

func (d *Provider) clean(name string) {
	if d.autoclean {
		if _, exist := d.factories[name]; exist {
			delete(d.factories, name)
		}
		if _, exist := d.defaultFactories[name]; exist {
			delete(d.defaultFactories, name)
		}
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
