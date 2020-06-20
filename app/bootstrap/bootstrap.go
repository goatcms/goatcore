package bootstrap

import (
	"fmt"
	"os"

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
func NewBootstrap(gapp app.App) *Bootstrap {
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
		appScope = b.gapp.AppScope()
		errs     []error
	)
	if !b.inited {
		return goaterr.Errorf("Bootstrap.Run must be run after modules init")
	}
	if b.runed {
		return goaterr.Errorf("Bootstrap.Run can not be run twice")
	}
	b.runed = true
	appScope.AddTasks(len(b.modules))
	for _, module := range b.modules {
		go func(module app.Module) {
			defer appScope.DoneTask()
			if err = module.Run(b.gapp); err != nil {
				errs = append(errs, err)
			}
		}(module)
	}
	if err = appScope.Wait(); err != nil {
		errs = append(errs, err)
	}
	return goaterr.ToError(goaterr.AppendError(errs, app.CloseApp(b.gapp)))
}

// ShowError print error / errors to stderr
func (b *Bootstrap) ShowError(inerr error) (code int, err error) {
	var (
		appScope = b.gapp.AppScope()
		deps     struct {
			ErrLVL string `argument:"errlvl"`
		}
		details   bool
		errorCode = 1
	)
	if err = appScope.InjectTo(&deps); err != nil {
		return errorCode, err
	}
	details = deps.ErrLVL == "details"
	if details {
		switch v := inerr.(type) {
		case goaterr.JSONError:
			fmt.Fprintf(os.Stderr, "\n%s", v.ErrorJSON())
		default:
			fmt.Fprintf(os.Stderr, "\n%s", v.Error())
		}
	} else {
		fmt.Fprintf(os.Stderr, "\n%s", inerr.Error())
	}
	return errorCode, nil
}
