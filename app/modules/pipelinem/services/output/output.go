package pip

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/pipelinem/services"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Output is piplines repository
type Output struct{}

// NewOutput create a Output instance
func NewOutput() *Output {
	return &Output{}
}

// OutputFactory create a Output instance
func OutputFactory(dp dependency.Provider) (ri interface{}, err error) {
	return services.PipOutput(NewOutput()), nil
}

// ScopeOutput return ScopeOutput instance for scope
func (output *Output) ScopeOutput(scope app.Scope) (output PipScopeOutput) {
	var (
		ctxScope = ctx.Scope()
		key      = pipPrefix + pip.Name
	)
	if _, err = ctxScope.Get(key); err == nil {
		return goaterr.Errorf("Output.Register: Script can not be define twice")
	}
	return ctxScope.Set(key, pip)
}
