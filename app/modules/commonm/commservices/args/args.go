package args

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/varutil/varg"
)

// ArgumentsProvider provider
type ArgumentsProvider struct {
	deps struct {
		StrictMode string `argument:"?strict"`
	}
	strictMode bool
}

// ArgumentsFactory create new Arguments instance
func ArgumentsFactory(dp app.DependencyProvider) (in interface{}, err error) {
	instance := &ArgumentsProvider{}
	if err = dp.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	if instance.strictMode, err = varg.MatchBool("StrictMode", instance.deps.StrictMode, false); err != nil {
		return nil, err
	}
	return commservices.Arguments(instance), nil
}

// StrictMode return strict mode switch
func (args *ArgumentsProvider) StrictMode() bool {
	return args.strictMode
}
