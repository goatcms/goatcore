package bootstrap

import (
	"fmt"

	"github.com/goatcms/goat-core/app"
)

// Bootstrap is default boot sequence
type Bootstrap struct {
	gapp    app.App
	modules []app.Module
	inited  bool
	runed   bool
}

// NewBootstrap create new Bootstrap object
func NewBootstrap(gapp app.App) app.Bootstrap {
	return &Bootstrap{
		gapp:    gapp,
		modules: []app.Module{},
		inited:  false,
		runed:   false,
	}
}

// Register add new module
func (b Bootstrap) Register(m app.Module) error {
	if b.inited {
		return fmt.Errorf("Can not add module after inited")
	}
	b.modules = append(b.modules, m)
	return nil
}

// Init all modules
func (b Bootstrap) Init() error {
	if b.inited {
		return fmt.Errorf("Bootstrap can not be inited twice")
	}
	b.inited = true
	for _, module := range b.modules {
		if err := module.RegisterDependencies(b.gapp); err != nil {
			return err
		}
	}
	for _, module := range b.modules {
		if err := module.InitDependencies(b.gapp); err != nil {
			return err
		}
	}
	// clean unused dependencies
	b.gapp = nil
	return nil
}

// Run all modules
func (b Bootstrap) Run() error {
	if !b.inited {
		return fmt.Errorf("Bootstrap.Run must be run after modules init")
	}
	if b.runed {
		return fmt.Errorf("Bootstrap.Run can not be run twice")
	}
	b.runed = true
	for _, module := range b.modules {
		if err := module.Run(); err != nil {
			return err
		}
	}
	return nil
}
