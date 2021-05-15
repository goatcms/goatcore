package datascope

import (
	"testing"

	"github.com/goatcms/goatcore/app"
)

func TestChildDataGetOverwritedValue(t *testing.T) {
	var (
		childScp  app.DataScope
		ivalue    interface{}
		ok        bool
		parentScp app.DataScope
		value     string
	)
	t.Parallel()
	parentScp = New(map[interface{}]interface{}{
		"overwritedKey": "value",
	})
	childScp = NewChild(parentScp, map[interface{}]interface{}{
		"overwritedKey": "overwrited-value",
	})
	ivalue = childScp.Value("overwritedKey")
	if value, ok = ivalue.(string); !ok {
		t.Errorf("Expected string and take: %v", ivalue)
		return
	}
	if value != "overwrited-value" {
		t.Errorf("expected value equals to 'overwrited-value' and take %s", value)
	}
}

func TestChildDataGetParentValue(t *testing.T) {
	var (
		childScp  app.DataScope
		ivalue    interface{}
		ok        bool
		parentScp app.DataScope
		value     string
	)
	t.Parallel()
	parentScp = New(map[interface{}]interface{}{
		"parentKey": "parent-value",
	})
	childScp = NewChild(parentScp, map[interface{}]interface{}{})
	ivalue = childScp.Value("parentKey")
	if value, ok = ivalue.(string); !ok {
		t.Errorf("Expected string and take: %v", ivalue)
		return
	}
	if value != "parent-value" {
		t.Errorf("expected value equals to 'parent-value' and take %s", value)
	}
}

func TestChildDataOverwriteParentValue(t *testing.T) {
	var (
		childScp  app.DataScope
		ivalue    interface{}
		ok        bool
		parentScp app.DataScope
		value     string
	)
	t.Parallel()
	parentScp = New(map[interface{}]interface{}{
		"key": "parent-value",
	})
	childScp = NewChild(parentScp, map[interface{}]interface{}{})
	ivalue = parentScp.Value("key")
	if value, ok = ivalue.(string); !ok {
		t.Errorf("Expected string and take: %v", ivalue)
		return
	}
	if value != "parent-value" {
		t.Errorf("Before overwrite expected value equals to 'parent-value' and take %s", value)
	}
	childScp.SetValue("key", "child-value")
	ivalue = childScp.Value("key")
	if value, ok = ivalue.(string); !ok {
		t.Errorf("Expected string and take: %v", ivalue)
		return
	}
	if value != "child-value" {
		t.Errorf("After overwrite expected value equals to 'child-value' and take %s", value)
	}
	// parent is not modified
	ivalue = parentScp.Value("key")
	if value, ok = ivalue.(string); !ok {
		t.Errorf("Expected string and take: %v", ivalue)
		return
	}
	if value != "parent-value" {
		t.Errorf("After overwrite expected parent value equals to 'parent-value' and take %s", value)
	}
}
