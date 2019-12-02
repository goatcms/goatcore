package bootstrap

import (
	"sync"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil/goaterr"
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
func (b *Bootstrap) Register(m app.Module) error {
	if b.inited {
		return goaterr.Errorf("Can not add module after inited")
	}
	b.modules = append(b.modules, m)
	return nil
}

// Init all modules
func (b *Bootstrap) Init() error {
	if b.inited {
		return goaterr.Errorf("Bootstrap can not be inited twice")
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
	return nil
}

// Run all modules
func (b *Bootstrap) Run() (err error) {
	var (
		waitGroup = &sync.WaitGroup{}
		errs      []error
	)
	if !b.inited {
		return goaterr.Errorf("Bootstrap.Run must be run after modules init")
	}
	if b.runed {
		return goaterr.Errorf("Bootstrap.Run can not be run twice")
	}
	b.runed = true
	waitGroup.Add(len(b.modules))
	for _, module := range b.modules {
		go func(module app.Module) {
			defer waitGroup.Done()
			if err = module.Run(b.gapp); err != nil {
				errs = append(errs, err)
			}
		}(module)
	}
	waitGroup.Wait()
	return goaterr.ToErrors(goaterr.AppendError(errs, app.CloseApp(b.gapp)))
}
