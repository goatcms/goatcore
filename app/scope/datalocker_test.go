package scope

import (
	"testing"

	"github.com/goatcms/goatcore/app"
)

func TestDataLocker(t *testing.T) {
	var (
		err    error
		value  string
		scp    app.DataScope
		locker app.DataScopeLocker
	)
	t.Parallel()
	scp = NewDataScope(make(map[string]interface{}))
	locker = scp.LockData()
	locker.Set("key", "value")
	if err = locker.Commit(); err != nil {
		t.Error(err)
		return
	}
	if value, err = GetString(scp, "key"); err != nil {
		t.Error(err)
		return
	}
	if value != "value" {
		t.Errorf("expected value equals to 'value' and take %s", value)
	}
}
