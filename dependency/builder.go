package dependency

import (
	"fmt"
)

// Builder is helper struct to build new instances
type Builder struct {
	Static    bool
	Factory   Factory
	Instance  Instance
	isCalled  bool
	isDefault bool
}

// Get build and return object instances
func (d *Builder) Get(dp Provider) (Instance, error) {
	var err error
	if d.isCalled == true {
		return nil, fmt.Errorf("Dependency detect circle. Circles are not allowed")
	}
	d.isCalled = true
	if d.Static == false {
		var instance Instance
		instance, err = d.factory(dp)
		if err != nil {
			d.isCalled = false
			return nil, err
		}
		d.isCalled = false
		return instance, nil
	}
	if d.Instance != nil {
		d.isCalled = false
		return d.Instance, nil
	}
	d.Instance, err = d.factory(dp)
	if err != nil {
		d.isCalled = false
		return nil, err
	}
	d.isCalled = false
	return d.Instance, nil
}

func (d *Builder) factory(dp Provider) (Instance, error) {
	var err error
	d.Instance, err = d.Factory(dp)
	if err != nil {
		return nil, err
	}
	return d.Instance, nil
}
