package datascope

import (
	"testing"

	"github.com/goatcms/goatcore/app"
)

func TestDataLocker(t *testing.T) {
	var (
		err    error
		ivalue interface{}
		value  string
		scp    app.DataScope
		locker app.DataScopeLocker
		ok     bool
	)
	t.Parallel()
	scp = New(make(map[interface{}]interface{}))
	locker = scp.LockData()
	locker.SetValue("key", "value")
	if err = locker.Commit(); err != nil {
		t.Error(err)
		return
	}
	ivalue = scp.Value("key")
	if value, ok = ivalue.(string); !ok {
		t.Errorf("Expected string and take: %v", ivalue)
		return
	}
	if value != "value" {
		t.Errorf("expected value equals to 'value' and take %s", value)
	}
}
