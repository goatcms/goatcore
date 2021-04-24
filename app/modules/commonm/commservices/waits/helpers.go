package waits

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// GetScopeWaitManager get ScopeWaitManager value from data scope
func GetScopeWaitManager(scp app.DataScope, name string) (value commservices.ScopeWaitManager, err error) {
	var (
		ins interface{}
		ok  bool
	)
	if value, ok = scp.Value(name).(commservices.ScopeWaitManager); !ok {
		return nil, goaterr.Errorf("%v %T is not a ScopeWaitManager", ins, ins)
	}
	return value, nil
}
