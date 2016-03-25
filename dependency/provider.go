package dependency

import (
	"fmt"
)

type DefaultProvider struct {
	pool map[string]*Builder
}

func NewProvider() Provider {
	return &DefaultProvider{
		pool: map[string]*Builder{},
	}
}

func (d *DefaultProvider) Get(faceName string) (Instance, error) {
	if row, exist := d.pool[faceName]; exist {
		instance, err := row.Get()
		if err != nil {
			return nil, err
		}
		return instance, nil
	}
	return nil, fmt.Errorf("implementation of interface ", faceName, " not exist")
}

func (d *DefaultProvider) AddService(name string, factory Factory) error {
	return d.add(name, factory, true)
}

func (d *DefaultProvider) AddFactory(name string, factory Factory) error {
	return d.add(name, factory, false)
}

func (d *DefaultProvider) add(name string, factory Factory, s bool) error {
	if _, exist := d.pool[name]; exist {
		return fmt.Errorf("interface ", name, " exist")
	}
	d.pool[name] = &Builder{
		Static:   s,
		Factory:  factory,
		Instance: nil,
	}
	return nil
}
