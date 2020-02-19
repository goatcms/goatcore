package scope

import (
	"testing"

	"github.com/goatcms/goatcore/app"
)

func TestChildDataGetOverwritedValue(t *testing.T) {
	var (
		err       error
		value     string
		parentScp app.DataScope
		childScp  app.DataScope
	)
	t.Parallel()
	parentScp = NewDataScope(map[string]interface{}{
		"overwritedKey": "value",
	})
	childScp = NewChildDataScope(parentScp, map[string]interface{}{
		"overwritedKey": "overwrited-value",
	})
	if value, err = GetString(childScp, "overwritedKey"); err != nil {
		t.Error(err)
		return
	}
	if value != "overwrited-value" {
		t.Errorf("expected value equals to 'overwrited-value' and take %s", value)
	}
}

func TestChildDataGetParentValue(t *testing.T) {
	var (
		err       error
		value     string
		parentScp app.DataScope
		childScp  app.DataScope
	)
	t.Parallel()
	parentScp = NewDataScope(map[string]interface{}{
		"parentKey": "parent-value",
	})
	childScp = NewChildDataScope(parentScp, map[string]interface{}{})
	if value, err = GetString(childScp, "parentKey"); err != nil {
		t.Error(err)
		return
	}
	if value != "parent-value" {
		t.Errorf("expected value equals to 'parent-value' and take %s", value)
	}
}

func TestChildDataOverwriteParentValue(t *testing.T) {
	var (
		err       error
		value     string
		parentScp app.DataScope
		childScp  app.DataScope
	)
	t.Parallel()
	parentScp = NewDataScope(map[string]interface{}{
		"key": "parent-value",
	})
	childScp = NewChildDataScope(parentScp, map[string]interface{}{})
	if value, err = GetString(childScp, "key"); err != nil {
		t.Error(err)
		return
	}
	if value != "parent-value" {
		t.Errorf("Before overwrite expected value equals to 'parent-value' and take %s", value)
	}
	childScp.Set("key", "child-value")
	if value, err = GetString(childScp, "key"); err != nil {
		t.Error(err)
		return
	}
	if value != "child-value" {
		t.Errorf("After overwrite expected value equals to 'child-value' and take %s", value)
	}
	// parent is not modified
	if value, err = GetString(parentScp, "key"); err != nil {
		t.Error(err)
		return
	}
	if value != "parent-value" {
		t.Errorf("After overwrite expected parent value equals to 'parent-value' and take %s", value)
	}
}
