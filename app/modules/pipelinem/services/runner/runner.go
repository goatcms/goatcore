package pip

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/pipelinem/services"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Runner is piplines repository
type Runner struct{}

// NewRunner create a Runner instance
func NewRunner() *Runner {
	return &Runner{}
}

// RunnerFactory create a Runner instance
func RunnerFactory(dp dependency.Provider) (ri interface{}, err error) {
	return services.PipRunner(NewRunner()), nil
}

// Register add/set new pip
func (runner *Runner) Register(ctx app.IOContext, pip services.Pip) (err error) {
	var (
		ctxScope = ctx.Scope()
		key      = pipPrefix + pip.Name
	)
	if _, err = ctxScope.Get(key); err == nil {
		return goaterr.Errorf("Runner.Register: Script can not be define twice")
	}
	return ctxScope.Set(key, pip)
}

// Get return pip by name
func (runner *Runner) Get(ctx app.IOContext, name string) (pip services.Pip, err error) {
	var (
		ctxScope = ctx.Scope()
		key      = pipPrefix + pip.Name
		pipi     interface{}
	)
	if pipi, err = ctxScope.Get(key); err != nil {
		return pip, goaterr.Wrapf("Runner.Get: Pip is undefines", err)
	}
	return pipi.(services.Pip), nil
}
